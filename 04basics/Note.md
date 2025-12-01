
    # Go Basics: Arrays, Default Values, and Unicode

    ## Default Values

    - `int` → `0`
    - `string` → `""` (empty string)
    - `bool` → `false`

    ---

    ## Array Declaration

    Declare and initialize an array in a single line:

    ```go
    nums := [3]int{1, 2, 3}
    fmt.Println(nums)
    ```

    ### 2D Arrays

    ```go
    nums := [2][2]int{{3, 4}, {5, 6}}
    fmt.Println(nums)
    ```

    **Array properties:**
    - Fixed size, predictable
    - Memory optimization
    - Constant time access

    ---

    ## Unicode and Runes in Go

    ```go
    for i, c := range "golang" {
            fmt.Println(i, string(c))
    }
    ```

    ### Explanation
    - **Unicode code point (rune):**
        - In Go, a string is a sequence of bytes, but `range` iterates over Unicode code points (runes), not bytes.
        - A rune is an alias for `int32` and represents a Unicode code point (a character).
    - **Starting byte of rune:**
        - The first value (`i`) in `for i, c := range ...` is the byte index where the current rune starts in the string.
        - For ASCII, each character is 1 byte, so the index increases by 1 each time.
        - For multi-byte (non-ASCII) characters, the index jumps by more than 1.
    - **Example: `300 -> 1 byte, 2 byte`:**
        - Some Unicode code points (characters) are represented with 1 byte (like ASCII), others need more bytes (like some non-English characters).
        - For example, code point 300 (U+012C, "Ĭ") is encoded in UTF-8 as 2 bytes.
