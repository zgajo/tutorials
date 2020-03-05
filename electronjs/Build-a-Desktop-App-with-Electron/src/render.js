const videoElement = document.querySelector("video");
const startBtn = document.getElementById("startBtn");
const stopBtn = document.getElementById("stopBtn");
const videoSelectBtn = document.getElementById("videoSelectBtn");

startBtn.onclick = e => {
  mediaRecorder.start();
  startBtn.classList.add("is-danger");
  startBtn.innerText = "Recording";
};

stopBtn.onclick = e => {
  mediaRecorder.stop();
  startBtn.classList.remove("is-danger");
  startBtn.innerText = "Start";
};

videoSelectBtn.onclick = getVideoSources;

const { desktopCapturer, remote } = require("electron");

const { dialog, Menu } = remote;

async function getVideoSources() {
  const inputSources = await desktopCapturer.getSources({
    types: ["window", "screen"]
  });

  const videoOptionsMenu = Menu.buildFromTemplate(
    inputSources.map(source => {
      return {
        label: source.name,
        click: () => selectSource(source)
      };
    })
  );

  videoOptionsMenu.popup();
}

let mediaRecorder;
const recorderChunks = [];

async function selectSource(source) {
  videoSelectBtn.innerText = source.name;

  const constraints = {
    video: {
      mandatory: {
        chromeMediaSource: "desktop",
        chromeMediaSourceId: source.id
      }
    }
  };

  const stream = await navigator.mediaDevices.getUserMedia(constraints);

  videoElement.srcObject = stream;

  videoElement.play();

  mediaRecorder = new MediaRecorder(stream, {
    mimeType: "video/webm; codecs=vp9"
  });

  mediaRecorder.ondataavailable = handleDataAvailable;
  mediaRecorder.onstop = handleStop;

  function handleDataAvailable(e) {
    console.log("video data available");
    recorderChunks.push(e.data);
  }

  const { writeFile } = require("fs");

  async function handleStop() {
    const blob = new Blob(recorderChunks, {
      mimeType: "video/webm; codecs=vp9"
    });

    const buffer = Buffer.from(await blob.arrayBuffer());

    const { filePath } = await dialog.showSaveDialog({
      buttonLabel: "Save video",
      defaultPath: `vid-${Date.now()}.webm`
    });

    console.log(filePath);

    writeFile(filePath, buffer, () => console.log("Your file has been saved"));
  }
}
