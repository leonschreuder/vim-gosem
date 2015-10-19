" Vim syntax file
" Language:	Go

" Save the cursor position to prevent jumping around
" let s:preRunWindowState = winsaveview()




function! HighlightFields()
ruby << EOR

cb = VIM::Buffer.current

# Calling go from shell for res
fields = `go run ~/go/src/leonmoll.de/vim-gosem/parser3.go -f #{cb.name}`

# puts "fields: " + fields

# adding fields to highlight group
# VIM.command( 'syn match goFields "\<' + fields + '\>"')
VIM.command( "syn keyword goFields #{fields}")

EOR
endfunction

call HighlightFields()


" highlight named group specified with 'syn keyword' or 'syn match'
hi def link     goFields         Function


" Match starting at the function definition til the first (
" This is the entire start of the line
" It ends with the entire endling line. Since the first
" occurrenc will mark the end of the highlight group

" syn region firstMethod start="^func Test_main(" end="^}"
" hi def link     firstMethod     Function





" Reset the cursor position to prevent jumping around
" call winrestview(s:preRunWindowState)



"--------------------------------------------------------------------------------
" Reference
"--------------------------------------------------------------------------------
" [1] search-replace trick to retreive stuff from the current buffer
"	http://stackoverflow.com/questions/9079561/how-to-extract-regex-matches-using-vim
"--------------------------------------------------------------------------------
