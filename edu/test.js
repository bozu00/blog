

var chunks = [];
const playbackVideo = document.getElementById('playback_video');
const playbackVideoSource = document.getElementById('playback_video_source');
var canv = document.querySelector('canvas');
var stream = canv.captureStream(60);
let blobUrl = null;
var recorder = null;
convertedStream = null; // deleted

var videoBlob = null;
var record_start_button = document.getElementById('record_start');
var record_stop_button = document.getElementById('record_stop');
navigator.getUserMedia = ( navigator.getUserMedia ||
        navigator.webkitGetUserMedia ||
        navigator.mozGetUserMedia ||
        navigator.msGetUserMedia);


  var img = new Image();
  img.src = "./math.jpg?" + new Date().getTime();
  console.log(img)
  /* 画像が読み込まれるのを待ってから処理を続行 */
  img.onload = function() {
    ctx.drawImage(img, 0, 0);
  }


window.onload = function () {
    navigator.getUserMedia(
            {audio : true},
            function(audioStream){
                //document.querySelector('audio').srcObject = stream;
                var audioContext = new AudioContext();
                audioContext.createMediaStreamSource(audioStream);

                // todo captureStreamをstopさせてない
                var canvasStream = canv.captureStream(60);
                video.srcObject = canvasStream;
                video.play();


                convertedStream = new MediaStream();
                convertedStream.addTrack(canvasStream.getVideoTracks()[0]);
                convertedStream.addTrack(audioStream.getAudioTracks()[0]);

                var options = {
                    audioBitsPerSecond : 128000,
                    videoBitsPerSecond : 2500000,
                    mimeType : 'video/webm'
                }
                recorder = new MediaRecorder(convertedStream, options);
                recorder.ondataavailable = function(e) {
                    chunks.push(e.data);
                };

                recorder.onstart = function(e) {
                    console.log("start")
                        recorder.requestData();
                };
                recorder.onstop = function(e) {
                    recorder = null;
                    playRecorded();
                };

            },
            console.log
                );

}

function record_save() {
    recorder.stop();
};

function playRecorded() {
    videoBlob = new Blob(chunks, { type: "video/webm" });
    blobUrl = window.URL.createObjectURL(videoBlob);

    var videofile = new File([videoBlob], "test.webm");
    localStorage['video'] = videofile;


    if (playbackVideo.src) {
        window.URL.revokeObjectURL(playbackVideo.src); // 解放
        playbackVideo.src = null;
    }

    playbackVideo.loadedmetadata = function(e) {
        console.log("loadedmetadata");
        console.log(e);
    }

    blobUrl = window.URL.createObjectURL(videofile);
    playbackVideo.src = blobUrl;

    playbackVideo.load();
    playbackVideo.play();


    console.group('metadata読み込み後');
    console.log('長さ:', playbackVideo.duration); 
    console.groupEnd();

    var link = document.createElement("a"); 
    link.href = blobUrl;
    link.download = "video.webm";
    link.innerHTML = "Click here to download the file";
    document.body.appendChild(link);


}

function record_start() {
    recorder.start(1000); 

  var img = new Image();
  img.src = "./math.jpg?" + new Date().getTime();
  console.log(img)
  /* 画像が読み込まれるのを待ってから処理を続行 */
  img.onload = function() {
    ctx.drawImage(img, 0, 0);
  }

}

function next_slide() {
  var img2 = new Image();
  img2.src = "./math2.png?" + new Date().getTime();
  img2.onload = function() {
    ctx.drawImage(img2, 0, 0);
  }
};

