

// Note: please restart the page if syntax highlighting works bad.
function isPalindrome2(word: string): boolean {
 const cleanedWord = word.replace(/[^a-zA-Z]/g, '').toLowerCase();
 return cleanedWord === cleanedWord.split('').reverse().join('');
}

function longest2(doc: string): string {
 const words = doc.split(/\b/);
 const palindromes = words.filter(word => isPalindrome(word));
 return palindromes.reduce((longest, current) => current.length > longest.length ? current : longest, '');
}



function isPalindrome(word: string): boolean {
  const n = word.length - 1;
  for (let i = 0; i <= n / 2; i++) {
    if (word[i] != word[n - i]) {
      return false;
    }
  }
  return true;
}

function longest(words: string[]): string {
  var longest = "";
  for (let i = 0; i < words.length; i++) {
    let word = words[i];
    if (longest.length < word.length && isPalindrome(words[i])) {
      longest = word;
    }
  }
  return longest;
}


const doc = "aaaa bb c racecar abcedfdecbx"
const words = doc.split(/\b/);
console.log(longest(words))
