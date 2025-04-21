let extract_classes_from_string content =
  let regex = Str.regexp "class=\\\"\\([^\\\"]+\\)\\\"" in
  let rec aux pos acc =
    try
      let _ = Str.search_forward regex content pos in
      let classes = Str.matched_group 1 content |> String.split_on_char ' ' in
      aux (Str.match_end ()) (classes @ acc)
    with Not_found -> acc
  in
  aux 0 [] |> List.sort_uniq String.compare


let extract_classes_from_files files =
  files
  |> List.map (fun filename ->
         let ic = open_in filename in
         let content = really_input_string ic (in_channel_length ic) in
         close_in ic;
         extract_classes_from_string content)
  |> List.flatten
  |> List.sort_uniq String.compare