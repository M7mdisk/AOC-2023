day=$(date +%-d)
name="day${day}"
token=$(cat .token)

mkdir $name
cd  $name
lynx --dump "https://adventofcode.com/2023/day/$day" | sed -n '/---/,/To play/p' > $name.txt
wget "https://adventofcode.com/2023/day/$day/input" -O input.txt --header="Cookie: session=$token"
