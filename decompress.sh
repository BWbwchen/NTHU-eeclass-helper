mkdir hw
find . -iname '*.zip' -exec sh -c 'mv {} hw/' ';'

# decompress the whole zip
find . -iname '*.zip' -exec sh -c 'unzip -o -d "${0%.*}" "$0"' '{}' ';'

# rename, only remain student id
for file in hw/*/*;
do
    mv "${file}" "${file% (*}"
done

# decompress each student file
find hw/*/* -iname '*.zip' -exec sh -c 'unzip -o -d "${0%.*}" "$0"' '{}' ';'
