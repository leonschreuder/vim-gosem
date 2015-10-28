" Vim syntax file
" Language:	Go

" Save the cursor position to prevent jumping around
" let s:preRunWindowState = winsaveview()


let s:script_folder_path = escape( expand( '<sfile>:p:h' ), '\' ) . '/'


function! HighlightFields()
ruby << EOR

cb = VIM::Buffer.current
cwd = VIM::evaluate("s:script_folder_path")

# Parsing the file
parserOut = `go run #{cwd}../parser3.go -f #{cb.name}`
groups = parserOut.split("|")
variableGroup = groups[0]
methodGroups = groups[1..-1]


# Highlighting the fields
if variableGroup.length > 0
    VIM.command( "syn keyword goFields #{variableGroup}")
end



# Highlighting the method variables
methodGroups.each { |methodGroup|
    methodGroupSplit = methodGroup.split(",")
    startLineNo = methodGroupSplit[0].to_i
    endLineNo = methodGroupSplit[1].to_i

    methodVariables = methodGroupSplit[2]
    startLine = cb[startLineNo].gsub(/\[|\]|\*/){|m|"\\" + m}
    endLine = cb[endLineNo]
    regionName = "method_on_line_" + startLineNo.to_s
    varGroupName = "vars_for_" + regionName


    # Highlighting variables as keywords
    VIM.command("syn keyword #{varGroupName} #{methodVariables} contained")


    # Making a region to contain the highlights
    VIM.command(
        "syn region #{regionName}" +
        " start=\"^#{startLine}\"" +
        " end=\"^#{endLine}\"" +
        " contains=#{varGroupName}" +
        (variableGroup.length > 0 ? ",goFields" : "") # Only add fields if we have some
    )
        

    VIM.command( "hi def link     #{varGroupName}     Statement")
}

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
" hi def link     goVars     Statement





" Reset the cursor position to prevent jumping around
" call winrestview(s:preRunWindowState)



"--------------------------------------------------------------------------------
" Reference
"--------------------------------------------------------------------------------
" [1] search-replace trick to retreive stuff from the current buffer
"	http://stackoverflow.com/questions/9079561/how-to-extract-regex-matches-using-vim
"--------------------------------------------------------------------------------
