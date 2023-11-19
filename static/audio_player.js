import { getAllAudioFiles } from "./scripts.js";

var audio           = document.getElementById('audioPlayer');
var volumeSlider    = document.getElementById('volumeSlider');
var seekBar         = document.getElementById('seekBar');

var currentTime     = document.getElementById('currentTime');
var totalTime       = document.getElementById('totalTime');

var startButton     = document.getElementById('startButton');
var pauseButton     = document.getElementById('pauseButton');
var stopButton      = document.getElementById('stopButton');

var backButton      = document.getElementById('backButton');
var forwardButton   = document.getElementById('forwardButton');

var nextButton      = document.getElementById('nextButton');

var currentTrack = 0;

var tracks = getAllAudioFiles();
console.log(tracks);

audio.addEventListener('loadedmetadata', function() {
    totalTime.textContent = formatTime(audio.duration);
    seekBar.max = audio.duration;
});

audio.addEventListener('timeupdate', function() {
    seekBar.value = audio.currentTime;
    currentTime.textContent = formatTime(audio.currentTime);
});

seekBar.addEventListener('input', function() {
    audio.currentTime = seekBar.value;
});

volumeSlider.addEventListener('input', function() {
    audio.volume = volumeSlider.value;
});

startButton.addEventListener('click', function() {
    audio.play();
});

pauseButton.addEventListener('click', function() {
    audio.pause();
});

stopButton.addEventListener('click', function() {
    audio.pause();
    audio.currentTime = 0;
});

backButton.addEventListener('click', function() {
    audio.currentTime -= 10;
});

forwardButton.addEventListener('click', function() {
    audio.currentTime += 10;
});

nextButton.addEventListener('click', function() {
    currentTrack += 1;
    if (currentTrack < tracks.length) {
        audio.src = "/static/" + tracks[currentTrack];
    }
});

function formatTime(time) {
    var minutes = Math.floor(time / 60);
    var seconds = Math.floor(time % 60);
    if (seconds < 10) {
        return minutes + ':0' + seconds;
    } else {
        return minutes + ':' + seconds;
    }
}

