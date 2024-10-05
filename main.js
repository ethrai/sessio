const configuration = {
  iceServers: [
    {
      urls: ["stun:stun1.l.google.com:19302", "stun:stun2.l.google.com:19302"],
    },
  ],
  iceCandidatePoolSize: 10,
};

let localStream = null;
let remoteStream = null;
let peerConnection = new RTCPeerConnection(configuration);
let roomDialog = null;
let roomId = null;

function init() {
  document
    .querySelector("#cameraOnBtn")
    .addEventListener("click", openUserMedia);
  
  document.querySelector("#createOfferBtn").addEventListener("click", createOffer);
  document.querySelector("#createAnswerBtn").addEventListener("click", createAnswer);
  document
    .querySelector("#addAnswerBtn")
    .addEventListener("click", addAnswer);
}

async function createOffer() {
  document.querySelector("#createOfferBtn").disabled = true;
  document.querySelector("#createAnswerBtn").disabled = true;

  peerConnection.onicecandidate = async (event) => {
    if (event.candidate) {
      document.querySelector("#offer").value = JSON.stringify(
        peerConnection.localDescription,
      );
    }
  };

  const offer = await peerConnection.createOffer();
  await peerConnection.setLocalDescription(offer);

  document.querySelector("#offer").value = JSON.stringify(offer);
}

async function createAnswer() {
  document.querySelector("#createOfferBtn").disabled = true;
  document.querySelector("#addAnswerBtn").disabled = true;

  peerConnection.onicecandidate = async (event) => {
    if (event.candidate) {
      document.querySelector("#answer").value = JSON.stringify(
        peerConnection.localDescription
      );
    }
  }

  let offer = document.querySelector("#offer").value;
  if (!offer) {
    alert("offer is empty!");
  }

  offer = JSON.parse(offer);
  await peerConnection.setRemoteDescription(offer);

  let answer = await peerConnection.createAnswer();
  await peerConnection.setLocalDescription(answer);
}

async function addAnswer() {
  let answer = document.querySelector("#answer").value;
  if (!answer) alert("no answer set!");

  answer = JSON.parse(answer);

  if (!peerConnection.currentRemoteDescription) {
    peerConnection.setRemoteDescription(answer);
  }
}

async function openUserMedia() {
  const stream = await navigator.mediaDevices.getUserMedia({
    video: true,
    audio: false,
  });
  document.querySelector("#localVideo").srcObject = stream;
  localStream = stream;
  remoteStream = new MediaStream();
  document.querySelector("#remoteVideo").srcObject = remoteStream;


    localStream.getTracks().forEach((track) => {
        peerConnection.addTrack(track, localStream);
    });

    peerConnection.ontrack = (event) => {
        event.streams[0].getTracks().forEach((track) => {
        remoteStream.addTrack(track);
        });
    };

  document.querySelector("#cameraOnBtn").disabled = true;
}

init();
