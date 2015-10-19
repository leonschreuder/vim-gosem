" Vim syntax file
" Language:	Java

" let s:preRunWindowState = winsaveview()

"--------------------------------------------------------------------------------
" Highlighting field names
"--------------------------------------------------------------------------------
function! HighlightFields()
python << EOF
import vim



# for result in allResults:

vim.command( 'syn match goFields "\<' + 'fmt' + '\>"')

EOF
endfunction

call HighlightFields()

" syn keyword     fset         call shell("")
hi def link     goFields         Function

" call winrestview(s:preRunWindowState)



"--------------------------------------------------------------------------------
" Reference
"--------------------------------------------------------------------------------
" [1] search-replace trick to retreive stuff from the current buffer
"	http://stackoverflow.com/questions/9079561/how-to-extract-regex-matches-using-vim
"--------------------------------------------------------------------------------
