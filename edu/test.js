

var chunks = [];
const playbackVideo = document.getElementById('playback_video');
const playbackModalVideo = document.getElementById('playback_modal_video');
const playbackVideoSource = document.getElementById('playback_video_source');
const screen = document.getElementById('screen');
var canv = document.querySelector('canvas');
var stream = canv.captureStream(60);
var blobUrl = null;
var recorder = null;

convertedStream = null; // deleted

var canvas_width = null;
var canvas_height = null;

var videoBlob = null;
var record_start_button = document.getElementById('record_start');
var record_stop_button = document.getElementById('record_stop');
navigator.getUserMedia = ( navigator.getUserMedia ||
        navigator.webkitGetUserMedia ||
        navigator.mozGetUserMedia ||
        navigator.msGetUserMedia);


var dragPosition = null;
var dropPosition = null;

class Slides {
    constructor(slides) {
        this.slides = slides;
    }
    sort(dragPosision, dropPosition) {
        console.log(`sort ${dragPosition} ${dropPosition}`);
        var drag_slide = this.slides[dragPosition];
        this.slides.splice(dropPosition, 0, drag_slide);
        var drag_index = dragPosition;
        if (dropPosition < dragPosition) {
            drag_index = drag_index + 1;
        }
        console.log(drag_index);
        this.slides.splice(drag_index, 1); //削除
    }
    addBySrc(src) {
        var new_slide = new Slide(src);
        this.slides.push(new_slide);
    }
}

class Slide {
    constructor(src) {
        this.src = src;
    }
}

s = new Slides( [ 
        new  Slide( "./idea1.jpg" ),
        new  Slide( "./idea2.jpg" ),
        new  Slide( "./idea3.jpg" ),
        new  Slide( "./idea4.jpg" ),
        new  Slide( "./idea5.jpg" ),
]);        

var vm = new Vue({
    el: '#slides',
    data: {
        slides: s,
        currenct_slide_index: 0
    },
    methods: {
        dragOverHandler: function(event) {
            console.log("over");
            event.currentTarget.classList.add('over');
            event.dataTransfer.dropEffect = 'move';
            event.stopPropagation();
            event.preventDefault(); //　必要!! これがないとdropイベントがハッカしない
        },
        dragStartHandler: function(event) {
            console.log("start");
            event.stopPropagation();
            dragPosition = Number(event.currentTarget.id);
            event.dataTransfer.setData('dummy', null); //必要!! これがないとfirefoxでdragoverがハッカしない

        },
        dragLeaveHandler: function(event) {
            //event.stopPropagation();
            console.log("leave");
            event.currentTarget.classList.remove('over');
            event.stopPropagation();
        },                            

        dropHandler: function(event) {
            console.log("drop");
            event.stopPropagation();
            dropPosition = Number(event.currentTarget.id);
            var tmp = event.dataTransfer.getData('dummy'); //必要!! これがないとfirefoxでdragoverがハッカしない
            vm.slides.sort(dragPosition, dropPosition);
            dragPosision = null;
            dropPosision = null;
            event.currentTarget.classList.remove("over");
        }

    }
});


function calculate_screen_size() {
    canvas_width  = screen.clientWidth;
    canvas_height = canvas_width / 16 * 9;
    if (canvas_height >= screen.clientHeight) {
        canvas_height = screen.clientHeight;
        canvas_width  = canvas_height / 9 * 16;
    }
    canv.height = canvas_height;
    canv.width  = canvas_width;
    playbackModalVideo.height = canvas_height;
    playbackModalVideo.width  = canvas_width;
}

window.onload = function () {


    console.log(screen.clientWidth);
    console.log(screen.clientHeight);
    calculate_screen_size();
    console.log(screen.clientWidth);
    console.log(screen.clientHeight);

    var img = new Image();
    img.src = "./idea.jpg?" + new Date().getTime();
    console.log(img)
        /* 画像が読み込まれるのを待ってから処理を続行 */
        img.onload = function() {
            ctx.drawImage(img, 0, 0, canvas_width, canvas_height);
        }

    var new_slide = document.querySelector('#new_slide');
    new_slide.addEventListener('dragover', newSlideDragOverHandler, true);
    new_slide.addEventListener('drop',     newSlideDropHandler, true);



    navigator.getUserMedia(
            {audio : true},
            function(audioStream){
                //document.querySelector('audio').srcObject = stream;
                var audioContext = new AudioContext();
                audioContext.createMediaStreamSource(audioStream);

                // todo captureStreamをstopさせてない
                var canvasStream = canv.captureStream(60);
                // video.srcObject = canvasStream;
                // video.play();

                convertedStream = new MediaStream();
                convertedStream.addTrack(canvasStream.getVideoTracks()[0]);
                convertedStream.addTrack(audioStream.getAudioTracks()[0]);

                var options = {
                    audioBitsPerSecond : 2500000,
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

    canv.style.display = "none"
    playbackModalVideo.style.display = "flex";

    if (playbackModalVideo.src) {
        window.URL.revokeObjectURL(playbackVideo.src); // 解放
        playbackVideo.src = null;
    }

    playbackModalVideo.loadedmetadata = function(e) {
        console.log("loadedmetadata");
        console.log(e);
    }

    blobUrl = window.URL.createObjectURL(videofile);
    playbackModalVideo.src = blobUrl;

    //一回おっきなcurrentTimeで再生してから、0から再生する (chromeのみだけど、再生時間が表示される)
    playbackModalVideo.currentTime = 1e101;
    playbackModalVideo.load();
    playbackModalVideo.play();
    playbackModalVideo.currentTime = 0;
    playbackModalVideo.play();


    console.group('metadata読み込み後');
    console.log('長さ:', playbackVideo.duration); 
    console.groupEnd();

    var link = document.createElement("a"); 
    link.href = blobUrl;
    link.download = "video.webm";
    link.innerHTML = "Click here to download the file";
    screen.appendChild(link);


}

function record_start() {
    recorder.start(1000); 

    var img = new Image();
    img.src = "./idea.jpg?" + new Date().getTime();
    console.log(img)
    /* 画像が読み込まれるのを待ってから処理を続行 */
    img.onload = function() {
        ctx.drawImage(img, 0, 0, canvas_width, canvas_height);
    }
}

function next_slide() {
    var next_image = new Image();
    vm.currenct_slide_index += 1;
    console.log(vm.currenct_slide_index);
    console.log(vm.slides.slides[vm.currenct_slide_index].src);
    next_image.src = vm.slides.slides[vm.currenct_slide_index].src + "?" + new Date().getTime();
    next_image.onload = function() {
        ctx.drawImage(next_image, 0, 0, canvas_width, canvas_height);
    }
};

function previous_slide() {
    var next_image = new Image();
    vm.currenct_slide_index -= 1;
    console.log(vm.currenct_slide_index);
    console.log(vm.slides.slides[vm.currenct_slide_index].src);
    next_image.src = vm.slides.slides[vm.currenct_slide_index].src + "?" + new Date().getTime();
    next_image.onload = function() {
        ctx.drawImage(next_image, 0, 0, canvas_width, canvas_height);
    }
};


function newSlideDragOverHandler(event) {
    event.preventDefault(); //　必要!! これがないとdropイベントがハッカしない
    event.stopPropagation();
}


function newSlideDropHandler(event) {
    event.preventDefault(); //　必要!! これがないとdropイベントがハッカしない
    event.stopPropagation();
    console.log("drop");
    files = event.dataTransfer.files;
    console.log(files);
    for (var i=0; i<files.length; i++) {
        if (!files[i] || files[i].type.indexOf('image/') < 0) {
            continue;
        }
        var fileReader = new FileReader();
        fileReader.onload = function( event ) {
            var loadedImageUri = event.target.result;
            console.log(loadedImageUri);
            vm.slides.addBySrc(loadedImageUri);
        };
        fileReader.readAsDataURL( files[i] );
    }
}
