" Vim syntax file
" Language:	Go

" Save the cursor position to prevent jumping around
" let s:preRunWindowState = winsaveview()

if exepath('gosem') == ""
    echo "Unable to highilght go source file, 'gosem' is not installed"
    finish
endif


" Comments; their contents
syn keyword     goTodo              contained TODO FIXME XXX BUG
syn cluster     goCommentGroup      contains=goTodo
syn region      goComment           start="/\*" end="\*/" contains=@goCommentGroup,@Spell
syn region      goComment           start="//" end="$" contains=@goCommentGroup,@Spell

hi def link     goComment           Comment
hi def link     goTodo              Todo

" Go escapes
syn match       goEscapeOctal       display contained "\\[0-7]\{3}"
syn match       goEscapeC           display contained +\\[abfnrtv\\'"]+
syn match       goEscapeX           display contained "\\x\x\{2}"
syn match       goEscapeU           display contained "\\u\x\{4}"
syn match       goEscapeBigU        display contained "\\U\x\{8}"
syn match       goEscapeError       display contained +\\[^0-7xuUabfnrtv\\'"]+

hi def link     goEscapeOctal       goSpecialString
hi def link     goEscapeC           goSpecialString
hi def link     goEscapeX           goSpecialString
hi def link     goEscapeU           goSpecialString
hi def link     goEscapeBigU        goSpecialString
hi def link     goSpecialString     Special
hi def link     goEscapeError       Error

syn match       goTestMethod        "\sTest_"
hi def link     goTestMethod        Todo
" hi def link     goTestMethod        Todo      " To close to variable
" hi def link     goTestMethod        Type      "Not strong enough
" hi def link     goTestMethod        Error     "To strong
" hi def link     goTestMethod        Underlined " to close to variable

" Strings and their contents
syn cluster     goStringGroup       contains=goEscapeOctal,goEscapeC,goEscapeX,goEscapeU,goEscapeBigU,goEscapeError
syn region      goString            start=+"+ skip=+\\\\\|\\"+ end=+"+ contains=@goStringGroup
syn region      goRawString         start=+`+ end=+`+

hi def link     goString            String
hi def link     goRawString         String

" Characters; their contents
syn cluster     goCharacterGroup    contains=goEscapeOctal,goEscapeC,goEscapeX,goEscapeU,goEscapeBigU
syn region      goCharacter         start=+'+ skip=+\\\\\|\\'+ end=+'+ contains=@goCharacterGroup

hi def link     goCharacter         Character

"--------------------------------------------------------------------------------
" Above is copied from go.vim
"--------------------------------------------------------------------------------



let s:script_folder_path = escape( expand( '<sfile>:p:h' ), '\' ) . '/'


function! HighlightFields()
ruby << EOR

cb = VIM::Buffer.current
cwd = VIM::evaluate("s:script_folder_path")

# Parsing the file
parserOut = `gosem -f #{cb.name}`
if parserOut.include?("|")

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
            " contains=" +
                "#{varGroupName}," +
                (variableGroup.length > 0 ? "goFields," : "") + # Only add fields if we have some
                "goString," +
                "goRawString," +
                "goComment," +
                "goTodo," +
                "goEscapeOctal," +
                "goEscapeC," +
                "goEscapeX," +
                "goEscapeU," +
                "goEscapeBigU," +
                "goSpecialString," +
                "goTestMethod," +
                "goEscapeError"
        )

        VIM.command( "hi def link     #{varGroupName}     Statement")
    }
end

EOR
endfunction
call HighlightFields()


" highlight named group specified with 'syn keyword' or 'syn match'
" hi def link     goFields         Todo
" hi def link     goFields         Type
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
