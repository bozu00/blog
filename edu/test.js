

var canv = document.querySelector('canvas');
//var video = document.querySelector('video');

// Optional frames per second argument.
var stream = canv.captureStream(60);
// Set the source of the <video> element to be the stream from the <canvas>.
video.srcObject = stream;
video.play();

recorder = new MediaRecorder(stream);

recorder.addEventListener('dataavailable', onRecordingReady);

function onRecordingReady(e) {
  var video = document.getElementById('video');
  video.srcObject = e.data;
  video.play();
}



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
