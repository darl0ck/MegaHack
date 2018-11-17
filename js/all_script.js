$(document).ready(function () {

    'use strict';
    setTimeout(()=>main(), 1000);
    function main() {
        const SpeechRecognition = window.SpeechRecognition || window.webkitSpeechRecognition;
        const SpeechGrammarList = window.SpeechGrammarList || window.webkitSpeechGrammarList;
        const SpeechRecognitionEvent = window.SpeechRecognitionEvent || window.webkitSpeechRecognitionEvent;
        const synth = window.speechSynthesis;
        const recognition = new SpeechRecognition();
        recognition.continuous = true;
        recognition.lang = 'ru-RU';
        recognition.interimResults = true;
        recognition.maxAlternatives = 1;
        recognition.start();

        recognition.onresult = (event) => {
            const res = (event.results[0][0].transcript);
            for (let el of event.results){
                console.info(el[0].transcript);
                if (el[0].transcript.toLowerCase().indexOf('главная')+1){
                    window.open('lk_main.html','_self');
                }else if(el[0].transcript.toLowerCase().indexOf('помощь')+1){
                    window.open('lk_help.html','_self');
                }else if(el[0].transcript.toLowerCase().indexOf('проекты')+1){
                    window.open('sign_up.html','_self');
                }else if(el[0].transcript.toLowerCase().indexOf('настройки')+1){
                    window.open('sign_up.html','_self');
                }
            }
        };
    }




    $(document).keydown(function (e) {
        if(e.keyCode === 32){
            e.preventDefault();
            window.location.replace("sign_up.html");
        }
    });


});
