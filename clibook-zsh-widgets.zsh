######################################
# Add note entry widget              #
# ####################################
add-clibook-entry-widget() {

  # Check if clibook is installed
  local required=("clibook")

  for tool in "${required[@]}"; do
    if ! command -v "$tool" &> /dev/null; then
      echo "Command '$tool' not installed or not in PATH" 
      zle -I
      return 1
    fi
  done

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
    clibook add --command "$last_command" --note $note
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

# Bind the widget to a key combination (e.g., Alt-E)
bindkey '^[t' add-clibook-entry-widget


######################################
# FZF search widget                  #
# ####################################
fzf-clibook-search-widget(){

  local required=("fzf" "clibook")

  for tool in "${required[@]}"; do
    if ! command -v "$tool" &> /dev/null; then
      echo "Command '$tool' not installed or not in PATH" 
      zle -I
      return 1
    fi
  done
  
  local selection=$(clibook --format csv | awk -F',' 'NR > 1 {print $1","$2}' | fzf -d, --preview 'clibook --id {1} --format csv | awk -F',' "NR > 1 {print \$3}"' | awk -F',' '{print $2}')

  BUFFER="$selection"
  CURSOR=${#BUFFER}

  zle reset-prompt
}

# Create a zsh widget for the function
zle -N fzf-clibook-search-widget

# Bind the widget to a key combination (e.g., Ctrl-T)
bindkey '^T' fzf-clibook-search-widget 
