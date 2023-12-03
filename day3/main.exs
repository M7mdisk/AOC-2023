
grid = File.read("input.txt") |> String.split("\n", trim: true) |> Enum.map(&String.split(&1, ""))  
digit_chars = ["1", "2", "3", "4", "5", "6", "7", "8", "9", "0"]
is_digit = &Enum.member?(digit_chars, &1)
is_symbol = &not Enum.member?(digit_chars ++ [".",""], &1)

neighbor_is_symbol = fn i, j ->
  rows = length(grid)
  cols = length(Enum.at(grid, 0))
  neighbors = [
    {i - 1, j},
    {i, j + 1},
    {i + 1, j},
    {i, j - 1},
    {i - 1, j - 1},
    {i - 1, j + 1},
    {i + 1, j - 1},
    {i + 1, j + 1}
  ]
  Enum.any?(neighbors, fn {x, y} ->
    x >= 0 and x < rows and y >= 0 and y < cols and is_symbol.(Enum.at(Enum.at(grid, x), y))
  end)
end

grid
|> Enum.with_index()
|> Enum.reduce({0,  true}, fn {row, i}, acc ->
  row
  |> Enum.with_index()
  |> Enum.reduce({elem(acc, 0), "", false}, fn {cell, j}, {total_so_far, cur_num, surrounded_by_digit} ->
    if is_digit.(cell) do 
      {total_so_far, cur_num <> cell, surrounded_by_digit || neighbor_is_symbol.(i, j)}
    else
        if surrounded_by_digit do
            if cur_num == "" do
              {total_so_far , "", false}
            else
              {total_so_far + String.to_integer(cur_num), "", false}
            end
        else
          {total_so_far, "", false}
        end
    end
  end)
end)
|> elem(0)
