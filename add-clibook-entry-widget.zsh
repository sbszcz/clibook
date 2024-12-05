# Define the widget function
add-clibook-entry-widget() {

 # Check if the environment variable is set
  if [[ -z "$CLIBOOK" ]]; then
    echo "Error: The environment variable CLIBOOK is not set. Please set CLIBOOK pointing to the clibook executable!"
    zle -I
    return 1
  fi

  # Define a temporary file
  local tmpfile=$(mktemp)
  
  # Get the last issued command
  local last_command=$(fc -ln -1)

  # Append the last issued command as information
  echo "last command: '$last_command'\n---\n" > $tmpfile

  # Open the editor in insert mode with cursor set below the informative text about the last issued command 
  ${EDITOR:-vim} +3 +'startinsert' "$tmpfile"

  if [[ -e "$tmpfile" ]]; then

    #Read the content of the temporary file and strip the informative text
    local note=$(< $tmpfile | awk '/---/{found=1; next} found')

    # Add entry into clibook
    ${CLIBOOK} add --command "$last_command" --note $note
    BUFFER=""
    CURSOR=${#BUFFER}
  fi

  # Clean up the temporary file
  rm -f "$tmpfile"

  # Refresh the zsh line editor
  zle reset-prompt
}

# Create a zsh widget for the function
zle -N add-clibook-entry-widget

# Bind the widget to a key combination (e.g., Ctrl-E)
bindkey '^E' add-clibook-entry-widget
