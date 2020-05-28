await Deno.symlink("./fixtures/symlink/file.txt", "./fixtures/symlink/symlink_file.txt");
const realPath = await Deno.realPath("./fixtures/symlink/file.txt");
const realSymLinkPath = await Deno.realPath("./fixtures/symlink/symlink_file.txt");

console.log(realPath);  // outputs "/home/alice/file.txt"
console.log(realSymLinkPath);  // outputs "/home/alice/file.txt"