$(document).ready(function () {

    'use strict';
    const SpeechRecognition = window.SpeechRecognition || window.webkitSpeechRecognition;
    const SpeechGrammarList = window.SpeechGrammarList || window.webkitSpeechGrammarList;
    const SpeechRecognitionEvent = window.SpeechRecognitionEvent || window.webkitSpeechRecognitionEvent;
    const synth = window.speechSynthesis;
    const field = document.querySelector('.speech-recognition__words');
    const microphoneIcon = document.querySelector('.speech-recognition__icon');
    let space_key = false;
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
            if (el[0].transcript.toLowerCase().indexOf('да')+1){
                window.open('/lk/main/','_self');
            }
        }
    };

    setTimeout(()=>main(), 1000);
    function main() {
        let txtToSay = new SpeechSynthesisUtterance('Нажмите пробел если хотите отключить голосовой помошник');
        synth.speak(txtToSay);
    }

    setTimeout(()=>saveHelp(), 9000);
    function saveHelp() {
        if(!space_key){
            let txtToSay = new SpeechSynthesisUtterance('У вас есть проблемы со зрением?');
            synth.speak(txtToSay);
        }
    }

    $(".btn_click").click(function () {
       var id = $(this).attr("id");
       $(".question").css("display","none");
       switch (id) {
           case "yes_1":
                $(".question_2").css("display","block");
               let txtToSay = new SpeechSynthesisUtterance('Использовать режим управления голосом?');
               synth.speak(txtToSay);
               break;

           case "no_1":
               $(".question_2").css("display","block");
               let txtToSay = new SpeechSynthesisUtterance('Использовать режим управления голосом?');
               synth.speak(txtToSay);
               break;
       }
    });

    $(document).keydown(function (e) {
        if(e.keyCode === 32){
            e.preventDefault();
            space_key = true
        }
    });

});
