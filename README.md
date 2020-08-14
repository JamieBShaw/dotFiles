# JBS - .Files

## Todo
- Make Swap file
- Install vim-plug
- Stow dotFiles
- install zsh-plugins from git


## Vim-plug

Run sh -c 'curl -fLo "${XDG_DATA_HOME:-$HOME/.local/share}"/nvim/site/autoload/plug.vim --create-dirs \
       https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim'
       
Above command runs in Alacritty not Konsole       
Check vim-plug path in init.vim 

## STOW

stow -t ~/.config/programmfolder -d . programmName


## Python 2 & 3 troubles with nvim

May require you to :checkhealth provider, if issues with python2 and 3 do below:

pip3 install pynvim

## Node

if error:
https://stackoverflow.com/questions/62873601/usr-local-bin-node-undefined-symbol-nghttp2-option-set-max-settings-erro-whe

install libnghttp2

sudo pacman -S libnghttp2


## zsh configuration:



### zsh plugins

#### zsh-syntax-highlighting
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git

### zsh-autosuggestions
git clone https://github.com/zsh-users/zsh-autosuggestions.git

### powerlevel10k
git clone --depth=1 https://github.com/romkatv/powerlevel10k.git

clone into ~/.config/zsh


