/**
 * @param {string} firstWord
 * @param {string} secondWord
 * @param {string} targetWord
 * @return {boolean}
 */
 var isSumEqual = function(firstWord, secondWord, targetWord) {
   const firstNumber = +convertNumber(firstWord)
   const secondNumber = +convertNumber(secondWord)
   const targetNumber = +convertNumber(targetWord)
   return firstNumber + secondNumber == targetNumber
};

function convertNumber(str){
    number = ''
    for (let i = 0; i < str.length; i++) {
        number+= str.charCodeAt(i) - 97
    }
    number = +number
    return number
}

console.log(isSumEqual("acb","cba","cdb"))