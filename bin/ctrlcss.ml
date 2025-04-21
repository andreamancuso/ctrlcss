open Ctrlcss_lib.Parser
open Ctrlcss_lib.Emitter

let () =
  let files = Sys.argv |> Array.to_list |> List.tl in
  let class_names = extract_classes_from_files files in
  let css = generate_css class_names in
  print_endline css