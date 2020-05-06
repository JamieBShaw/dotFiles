
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
#source ~/.zsh/zsh-autosuggestions/zsh-autosuggestions.zsh
source ~/.config/.zsh/zsh-autosuggestions/zsh-autosuggestions.zsh
source ~/.config/zsh-syntax-highlighting/zsh-syntax-highlighting.zsh
source ~/powerlevel10k/powerlevel10k.zsh-theme
fpath=(~/.config/.zsh/completion $fpath)
neofetch
# The following lines were added by compinstall

zstyle ':completion:*' completer _complete _ignored
zstyle :compinstall filename '/home/james/.zshrc'

bindkey -v
# End of lines added by compinstall
