" Vim syntax file
" Language:	Go

" Save the cursor position to prevent jumping around
" let s:preRunWindowState = winsaveview()


let s:script_folder_path = escape( expand( '<sfile>:p:h' ), '\' ) . '/'


function! HighlightFields()
ruby << EOR

cb = VIM::Buffer.current

# Calling go from shell for res

cwd = VIM::evaluate("s:script_folder_path")

parserOut = `go run C:/progs_manual/cygwin64#{cwd}../parser3.go -f #{cb.name}`
puts parserOut
#splitResult = parserOut.split("|")
#fields = splitResult[0]
#
#methodSplit = splitResult[1].split(",")
#
#startLineNo = methodSplit[0]
#endLineNo = methodSplit[1]
#vars = methodSplit[2]
#
## adding fields to highlight group
## VIM.command( 'syn match goFields "\<' + fields + '\>"')
#VIM.command( "syn keyword goFields #{fields}")
#
#VIM.command("syn region firstMethod start=\"^#{cb[startLineNo]}\" end=\"^#{cb[endLineNo]}\""

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
hi def link     firstMethod     Function





" Reset the cursor position to prevent jumping around
" call winrestview(s:preRunWindowState)



"--------------------------------------------------------------------------------
" Reference
"--------------------------------------------------------------------------------
" [1] search-replace trick to retreive stuff from the current buffer
"	http://stackoverflow.com/questions/9079561/how-to-extract-regex-matches-using-vim
"--------------------------------------------------------------------------------
