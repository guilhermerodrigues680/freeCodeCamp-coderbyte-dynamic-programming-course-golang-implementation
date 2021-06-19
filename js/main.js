/** @type{HTMLTextAreaElement} */
const textAreaStd = document.getElementById('textarea-std');

// stringUrl relativo ao arquivo que chamou o `main.js`
const myWorker = new Worker("js/worker.js");
console.log(myWorker);

myWorker.postMessage({ type: 'init', value: null })

myWorker.onmessage = function (ev) {
  console.debug('onmessage', ev);
  const message = ev.data;
  if (message?.type) {
    switch (message.type) {
    case 'write_to_stdout':
      updateTextStdout(message.message);
      break;
    case 'webassembly_started':
      document.getElementById("runButton").disabled = false;
      break;
    case 'webassembly_executed':
      document.getElementById("runButton").disabled = false;
      break;
    default:
      console.debug(`Mensagem ${message.type} desconhecida.`);
      break;
    }
  } else {
    console.debug(`Mensagem ${message} não possui tipo.`);
  }
};

myWorker.onerror = function (ev) {
  console.debug('onerror', ev);
};

myWorker.onmessageerror = function (ev) {
  console.debug('onmessageerror', ev);
};

function run() {
  myWorker.postMessage({ type: "run" });
  textAreaStd.value = "-> Rodando programa...\n";
};

function updateTextStdout(message) {
  textAreaStd.value += message + "\n";
};
