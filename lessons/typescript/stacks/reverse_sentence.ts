
// For more details, see: [reverse_sentence.md](https://github.com/boston-code-coffee/workshop/blob/main/lessons/problems/stacks/reverse_sentence.md)
function reverse(sentence: string): string {
  let fwd: string[] = sentence.split(' ');
  let rev : string[] = [];
  while(0 < fwd.length) {
     rev.push(fwd.pop())
  }
  return rev.join(' ')
}

console.log(reverse('hello is this correct'));

