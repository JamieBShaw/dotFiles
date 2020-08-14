export EDITOR="nvim"
HISTFILE=~/.histfile
HISTSIZE=2000
HISTSIZE=2000
SAVEHIST=2000
bindkey -e
zstyle :compinstall filename '/home/jamie/.zshrc'
POWERLEVEL9K_LEFT_PROMPT_ELEMENTS=(user dir vcs)
POWERLEVEL9K_RIGHT_PROMPT_ELEMENTS=(root_indicator)
POWERLEVEL9K_SHORTEN_DIR_LENGTH=1
autoload -Uz compinit && compinit -i
export ANDROID_HOME=$HOME/Library/Android/sdk
export PATH=$PATH:$ANDROID_HOME/emulator
export PATH=$PATH:$ANDROID_HOME/tools
export PATH=$PATH:$ANDROID_HOME/tools/bin
export PATH=$PATH:$ANDROID_HOME/platform-tools
source ~/.config/.zsh/zsh-autosuggestions/zsh-autosuggestions.zsh
source ~/.config/zsh-syntax-highlighting/zsh-syntax-highlighting.zsh
source ~/powerlevel10k/powerlevel10k.zsh-theme
fpath=(~/.config/.zsh/completion $fpath)
neofetch


# The following lines were added by compinstall
zstyle ':completion:*' completer _complete _ignored
zstyle :compinstall filename '/home/james/.zshrc'
alias -g vim="nvim"
alias -g vi="nvim"
alias -g v="nvim"
bindkey -v
eval "$(gh completion -s zsh)"
# End of lines added by compinstall
if [ -n "${NVIM_LISTEN_ADDRESS+x}" ]; then
  export COLORTERM="truecolor"
fi
[ -f ~/.fzf.zsh ] && source ~/.fzf.zsh
