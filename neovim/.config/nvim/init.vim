let g:fzf_install = 'yes | ./install'
call plug#begin('~/loc:al/share/nvim/plugged')

Plug 'neoclide/coc.nvim', {'branch': 'release'}
Plug 'vim-airline/vim-airline'
Plug 'vim-airline/vim-airline-themes'
Plug 'junegunn/fzf', { 'do': g:fzf_install }
Plug 'junegunn/fzf.vim'
Plug 'jremmen/vim-ripgrep'
Plug 'unblevable/quick-scope'
Plug 'airblade/vim-gitgutter'
Plug 'leafgarland/typescript-vim'
Plug 'pangloss/vim-javascript'
Plug 'ianks/vim-tsx'
Plug 'gruvbox-community/gruvbox'
Plug 'Yggdroot/indentLine'
Plug 'ap/vim-css-color' "Displays a preview of colors with CS
Plug 'luochen1990/rainbow'
Plug 'frazrepo/vim-rainbow'
Plug 'mbbill/undotree'
Plug 'scrooloose/nerdtree'
Plug 'scrooloose/nerdcommenter'
Plug 'Xuyuanp/nerdtree-git-plugin'
Plug 'tpope/vim-fugitive'
Plug 'tiagofumo/vim-nerdtree-syntax-highlight'
Plug 'ryanoasis/vim-devicons'
Plug 'pacha/vem-tabline'
Plug 'neomake/neomake'
Plug 'HerringtonDarkholme/yats.vim'
Plug 'fatih/vim-go', { 'do': ':GoUpdateBinaries' }
Plug 'sheerun/vim-polyglot'
call plug#end()

nnoremap <up> <nop>
nnoremap <down> <nop>
nnoremap <left> <nop>
nnoremap <right> <nop>

let mapleader=" "
syntax on
filetype plugin on

set hidden
set noerrorbells
set tabstop=2 softtabstop=2
set shiftwidth=2
set expandtab
set smartindent
set relativenumber
set mouse=a
set nu
set nowrap
set smartcase
set ignorecase
set noswapfile
set nobackup
set nowritebackup
set updatetime=300
set clipboard=unnamedplus
set undodir=~/.vim/undodir
set undofile
set incsearch
set nohlsearch
set cursorline
set wildmode=longest,list,full
set wildmenu
set wildignore+=**/node_modules/**
set termguicolors
set smarttab

if exists('+termguicolors')
    let &t_8f = "\<Esc>[38;2;%lu;%lu;%lum"
    let &t_8b = "\<Esc>[48;2;%lu;%lu;%lum"
endif

colorscheme gruvbox
set background=dark

highlight CursorLine cterm=NONE 

" --- vim go (polyglot) settings.
let g:go_highlight_build_constraints = 1
let g:go_highlight_extra_types = 1
let g:go_highlight_fields = 1
let g:go_highlight_functions = 1
let g:go_highlight_methods = 1
let g:go_highlight_operators = 1
let g:go_highlight_structs = 1
let g:go_highlight_types = 1
let g:go_highlight_function_parameters = 1
let g:go_highlight_function_calls = 1

let g:go_highlight_generate_tags = 1
let g:go_highlight_format_strings = 1
let g:go_highlight_variable_declarations = 1
let g:go_auto_sameids = 1

"Quickscopre highlight
let g:qs_highlight_on_keys = ['f', 'F', 't', 'T']
highlight QuickScopePrimary cterm=underline

" Vim-Go

let g:go_auto_type_info = 1
let g:go_fmt_command = "goimports"


let g:javascript_plugin_jsdoc = 1
let g:javascript_plugin_ngdoc = 1

nmap <C-n> :NERDTreeToggle<CR>
vmap ++ <plug>NERDCommenterToggle
nmap ++ <plug>NERDCommenterToggle
let NERDTreeShowHidden=1

let g:coc_global_extensions = [
  \ 'coc-snippets',
  \ 'coc-pairs',
  \ 'coc-yaml',
  \ 'coc-tsserver',
  \ 'coc-eslint', 
  \ 'coc-prettier', 
  \ 'coc-html',
  \ 'coc-json',
  \ 'coc-java',
  \ 'coc-python',
  \ 'coc-highlight',
  \ 'coc-markdownlint'
  \ ]

if executable('rg')
    let g:rg_derive_root='true'
endif

"Airline
let g:airline_powerline_fonts = 1

"Vim Go
autocmd BufWritePre *.go :call CocAction('runCommand', 'editor.action.organizeImport')


"Vim Rainbow
let g:rainbow_active = 1

let g:vrfr_rg = 'true'
let g:netrw_browse_split = 2
let g:netrw_banner = 0
let g:netrw_winsize = 25

nnoremap <leader>h :wincmd h<CR>
nnoremap <leader>j :wincmd j<CR>
nnoremap <leader>k :wincmd k<CR>
nnoremap <leader>l :wincmd l<CR>
nnoremap <leader>u :UndotreeShow<CR>
nnoremap <leader>pv :wincmd v<bar> :Ex <bar> :vertical resize 30<CR>
nnoremap <Leader>ps :Rg<SPACE>
nnoremap <silent> <Leader>+ :vertical resize +5<CR>
nnoremap <silent> <Leader>- :vertical resize -5<CR>

"FZF 
nnoremap <C-p> :FZF<cr>
nnoremap <C-f> :GFiles<CR>
nnoremap <silent> <leader>F :FZF ~<cr>
let g:fzf_layout = { 'window': { 'width': 0.8, 'height': 0.8 } }


"vem-tabline

nmap <leader>x :bdelete<CR>
nmap <leader>, <Plug>vem_prev_buffer-
nmap <leader>. <Plug>vem_next_buffer-
nmap <leader>1 :VemTablineGo 1<CR>
nmap <leader>2 :VemTablineGo 2<CR>
nmap <leader>3 :VemTablineGo 3<CR>
nmap <leader>4 :VemTablineGo 4<CR>
nmap <leader>5 :VemTablineGo 5<CR>
nmap <leader>6 :VemTablineGo 6<CR>
nmap <leader>7 :VemTablineGo 7<CR>
nmap <leader>8 :VemTablineGo 8<CR>
nmap <leader>9 :VemTablineGo 9<CR>

let g:vem_tabline_show = 2
let g:vem_tabline_show_number = 'index'
let g:vem_tabline_show_icon = 1



set updatetime=300

" don't give |ins-completion-menu| messages.
set shortmess+=c

" always show signcolumns
set signcolumn=yes

" Use tab for trigger completion with characters ahead and navigate.
" Use command ':verbose imap <tab>' to make sure tab is not mapped by other plugin.
inoremap <silent><expr> <TAB>
      \ pumvisible() ? "\<C-n>" :
      \ <SID>check_back_space() ? "\<TAB>" :
      \ coc#refresh()
inoremap <expr><S-TAB> pumvisible() ? "\<C-p>" : "\<C-h>"

function! s:check_back_space() abort
  let col = col('.') - 1
  return !col || getline('.')[col - 1]  =~# '\s'
endfunction

" Use <c-space> to trigger completion.
inoremap <silent><expr> <c-space> coc#refresh()

" Use <cr> to confirm completion, `<C-g>u` means break undo chain at current position.
" Coc only does snippet and additional edit on confirm.
" Or use `complete_info` if your vim support it, like:
" inoremap <expr> <cr> complete_info()["selected"] != "-1" ? "\<C-y>" : "\<C-g>u\<CR>"

" Use `[g` and `]g` to navigate diagnostics
nmap <silent> [g <Plug>(coc-diagnostic-prev)
nmap <silent> ]g <Plug>(coc-diagnostic-next)

" Remap keys for gotos
nmap <silent> gd <Plug>(coc-definition)
nmap <silent> gy <Plug>(coc-type-definition)
nmap <silent> gi <Plug>(coc-implementation)
nmap <silent> gr <Plug>(coc-references)

" Use K to show documentation in preview window
nnoremap <silent> D :call <SID>show_documentation()<CR>

function! s:show_documentation()
  if (index(['vim','help'], &filetype) >= 0)
    execute 'h '.expand('<cword>')
  else
    call CocAction('doHover')
  endif
endfunction

" Highlight symbol under cursor on CursorHold
autocmd CursorHold * silent call CocActionAsync('highlight')

" Remap for rename current word
nmap <F2> <Plug>(coc-rename)

" Remap for format selected region
xmap <leader>f  <Plug>(coc-format-selected)
nmap <leader>f  <Plug>(coc-format-selected)

augroup mygroup
  autocmd!
  " Setup formatexpr specified filetype(s).
  autocmd FileType typescript,json setl formatexpr=CocAction('formatSelected')
  " Update signature help on jump placeholder
  autocmd User CocJumpPlaceholder call CocActionAsync('showSignatureHelp')
augroup end

" Remap for do codeAction of selected region, ex: `<leader>aap` for current paragraph
xmap <leader>a  <Plug>(coc-codeaction-selected)
nmap <leader>a  <Plug>(coc-codeaction-selected)

" Remap for do codeAction of current line
nmap <leader>ac  <Plug>(coc-codeaction)
" Fix autofix problem of current line
nmap <leader>qf  <Plug>(coc-fix-current)

" Create mappings for function text object, requires document symbols feature of languageserver.
xmap if <Plug>(coc-funcobj-i)
xmap af <Plug>(coc-funcobj-a)
omap if <Plug>(coc-funcobj-i)
omap af <Plug>(coc-funcobj-a)

" Use <C-d> for select selections ranges, needs server support, like: coc-tsserver, coc-python
nmap <silent> <C-d> <Plug>(coc-range-select)
xmap <silent> <C-d> <Plug>(coc-range-select)

" Use `:Format` to format current buffer
command! -nargs=0 Format :call CocAction('format')

" Use `:Fold` to fold current buffer
command! -nargs=? Fold :call     CocAction('fold', <f-args>)

" use `:OR` for organize import of current buffer
command! -nargs=0 OR   :call     CocAction('runCommand', 'editor.action.organizeImport')

" Add status line support, for integration with other plugin, checkout `:h coc-status`
set statusline^=%{coc#status()}%{get(b:,'coc_current_function','')}

fun! TrimWhitespace()
    let l:save = winsaveview()
    keeppatterns %s/\s\+$//e
    call winrestview(l:save)
endfun
" Using CocList
" Show all diagnostics
nnoremap <silent> <space>a  :<C-u>CocList diagnostics<cr>
" Manage extensions
nnoremap <silent> <space>e  :<C-u>CocList extensions<cr>
" Show commands
nnoremap <silent> <space>c  :<C-u>CocList commands<cr>
" Find symbol of current document
nnoremap <silent> <space>o  :<C-u>CocList outline<cr>
" Search workspace symbols
nnoremap <silent> <space>s  :<C-u>CocList -I symbols<cr>
" Do default action for next item.
" Do default action for previous item.
" Resume latest coc list
nnoremap <silent> <space>p  :<C-u>CocListResume<CR>
