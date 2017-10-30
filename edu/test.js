

var chunks = [];
const playbackVideo = document.getElementById('playback_video');
var canv = document.querySelector('canvas');
var stream = canv.captureStream(60);
let blobUrl = null;
var recorder = null;


var record_start_button = document.getElementById('record_start');
var record_stop_button = document.getElementById('record_stop');
navigator.getUserMedia = ( navigator.getUserMedia ||
                       navigator.webkitGetUserMedia ||
                       navigator.mozGetUserMedia ||
                       navigator.msGetUserMedia);



window.onload = function () {
    navigator.getUserMedia(
            {audio : true},
            function(stream){
                // document.querySelector('audio').src = URL.createObjectURL(stream);
                document.querySelector('audio').srcObject = stream;
                var audioContext = new AudioContext();
                audioContext.createMediaStreamSource(stream);


                var canvasStream = canv.captureStream(60);
                video.srcObject = canvasStream;
                video.play();

                // recorder = new MediaRecorder(canvasStream);

                convertedStream = new MediaStream();
                convertedStream.addTrack(canvasStream.getVideoTracks()[0]);
                convertedStream.addTrack(stream.getAudioTracks()[0]);


                 var options = {
                     audioBitsPerSecond : 128000,
                     videoBitsPerSecond : 2500000,
                     mimeType : 'video/webm'
                 }
                
                recorder = new MediaRecorder(convertedStream);

                // recorder.addEventListener('dataavailable', onRecordingReady);
                recorder.ondataavailable = function(e) {
                    chunks.push(e.data);
                    console.log(e.data.type);
                };

                recorder.onstop = function(e) {
                    recorder = null;
                    playRecorded();
                };


                recorder.addEventListener('dataavailable', onRecordingReady);


            },
            console.log
                );

}

function record_save() {
    recorder.stop();
};

function playRecorded() {
    const videoBlob = new Blob(chunks, { type: "video/webm" });
    blobUrl = window.URL.createObjectURL(videoBlob);
    console.log(blobUrl);

    playbackVideo.src = blobUrl;
    //playbackVideo.srcObject = videoBlob;
    playbackVideo.play();

    if (playbackVideo.src) {
        window.URL.revokeObjectURL(playbackVideo.src); // 解放
        playbackVideo.src = null;
    }

    var link = document.createElement("a"); // Or maybe get it from the current document
    link.href = blobUrl;
    link.download = "video.webm";
    link.innerHTML = "Click here to download the file";
    document.body.appendChild(link);


}

function onRecordingReady(e) {
    chunks.push(e.data);
    console.log(chunks.length);
}

function record_start() {
    recorder.start(1000); 
}


/*`
recordButton = document.getElementById('record_start');
stopButton = document.getElementById('record_stop');

var recordButton, stopButton, recorder, liveStream;
navigator.mediaDevices.getUserMedia({
  audio: true,
  //video: true
})
.then(function (stream) {
  recordButton.disabled = false;
  recordButton.addEventListener('click', startRecording);
  stopButton.addEventListener('click', stopRecording);

});



function startRecording() {
  recorder = new MediaRecorder(liveStream);

  recorder.addEventListener('dataavailable', onRecordingReady);

  recordButton.disabled = true;
  stopButton.disabled = false;

  recorder.start();
}

function stopRecording() {
  recordButton.disabled = false;
  stopButton.disabled = true;

  recorder.stop();
}

function onRecordingReady(e) {
  var video = document.getElementById('recording');
  // e.data contains a blob representing the recording
  video.src = URL.createObjectURL(e.data);
  video.play();
}

*/




/*
 *
 *
 *
 *
 *


var recordButton, stopButton, recorder, liveStream;

window.onload = function () {
  recordButton = document.getElementById('record_start');
  stopButton = document.getElementById('record_stop');

  // get video & audio stream from user
  navigator.mediaDevices.getUserMedia({
    audio: true,
    //video: true
  })
  .then(function (stream) {
    recordButton.disabled = false;
    recordButton.addEventListener('click', startRecording);
    stopButton.addEventListener('click', stopRecording);

  });
};

function startRecording() {
  recorder = new MediaRecorder(liveStream);

  recorder.addEventListener('dataavailable', onRecordingReady);

  recordButton.disabled = true;
  stopButton.disabled = false;

  recorder.start();
}

function stopRecording() {
  recordButton.disabled = false;
  stopButton.disabled = true;

  recorder.stop();
}

function onRecordingReady(e) {
  var video = document.getElementById('recording');
  // e.data contains a blob representing the recording
  video.src = URL.createObjectURL(e.data);
  video.play();
}




*/












/*
var recordButton, stopButton, recorder, liveStream;

window.onload = function () {
  recordButton = document.getElementById('record');
  stopButton = document.getElementById('stop');

  // get video & audio stream from user
  navigator.mediaDevices.getUserMedia({
    audio: true
  })
  .then(function (stream) {
    var liveVideo = document.getElementById('live');
    liveVideo.src = URL.createObjectURL(stream);
    liveVideo.play();

    recordButton.disabled = false;
    recordButton.addEventListener('click', startRecording);
    stopButton.addEventListener('click', stopRecording);

  });
};

function startRecording() {
  recorder = new MediaRecorder(liveStream);

  recorder.addEventListener('dataavailable', onRecordingReady);

  recordButton.disabled = true;
  stopButton.disabled = false;

  recorder.start();
}

function stopRecording() {
  recordButton.disabled = false;
  stopButton.disabled = true;

  // Stopping the recorder will eventually trigger the 'dataavailable' event and we can complete the recording process
  recorder.stop();
}

function onRecordingReady(e) {
  var video = document.getElementById('recording');
  // e.data contains a blob representing the recording
  video.src = URL.createObjectURL(e.data);
  video.play();
}
*/
