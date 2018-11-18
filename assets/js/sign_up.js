$(document).ready(function () {

    'use strict';
    const SpeechRecognition = window.SpeechRecognition || window.webkitSpeechRecognition;
    const SpeechGrammarList = window.SpeechGrammarList || window.webkitSpeechGrammarList;
    const SpeechRecognitionEvent = window.SpeechRecognitionEvent || window.webkitSpeechRecognitionEvent;
    const synth = window.speechSynthesis;
    const field = document.querySelector('.speech-recognition__words');
    const microphoneIcon = document.querySelector('.speech-recognition__icon');
    let space_key = false;

    setTimeout(()=>main(), 1000);
    function main() {
        let txtToSay = new SpeechSynthesisUtterance('Нажмите пробел если хотите отключить голосовой помошник');
        synth.speak(txtToSay);
    }

    setTimeout(()=>saveHelp(), 9000);
    function saveHelp() {
        if(!space_key){
            let txtToSay = new SpeechSynthesisUtterance('Вы видите, что изображено на экране?');
            synth.speak(txtToSay);
        }
    }

    $(document).keydown(function (e) {
        if(e.keyCode === 32){
            e.preventDefault();
            space_key = true
        }
    });

});
