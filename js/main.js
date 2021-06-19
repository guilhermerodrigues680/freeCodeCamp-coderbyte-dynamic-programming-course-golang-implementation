//---- WORKER --------------------------------------------------------

// stringUrl relativo ao arquivo que chamou o `main.js`
const wasmWorker = new Worker("js/worker.js");
// wasmWorker.postMessage({ type: 'init', value: null })

wasmWorker.onmessage = function (ev) {
  // console.debug('onmessage', ev);
  const message = ev.data;
  if (message?.type) {
    switch (message.type) {
    case 'write_to_stdout':
      if (message.message) {
        updateWebassemblyOutput(message.message);
      } else {
        updateWebassemblyOutput(message.error, "stderr");
      }
      break;
    case 'webassembly_started':
      console.info(message.message);
      break;
    case 'running_webassembly':
      console.info(message.message);
      disableRunButton();
      break;
    case 'webassembly_terminated_with_error':
      console.info(message.error);
      enableRunButton();
      break;
    case 'webassembly_completed_successfully':
      console.info(message.message);
      enableRunButton();
      break;
    default:
      console.debug(`Mensagem ${message.type} desconhecida.`);
      break;
    }
  } else {
    console.debug(`Mensagem ${message} não possui tipo.`);
  }
};

wasmWorker.onerror = function (ev) {
  console.debug('onerror', ev);
};

wasmWorker.onmessageerror = function (ev) {
  console.debug('onmessageerror', ev);
};

//--------------------------------------------------------------------

const programSelector = twoWayDataBinding(document.getElementById('program-selector'));
const textAreaStd = twoWayDataBinding(document.getElementById('textarea-std'));
const runButton = document.getElementById('runButton');
const cancelRunButton = document.getElementById('cancel-run-button');

function twoWayDataBinding(el) {
  // 2 way data binding
  return {
    set value(v){ v === undefined || v === null || v === "" ? el.value = "" : el.value = v; },
    get value(){ return el.value; }
  }
}

function enableRunButton() {
  runButton.disabled = false;
  runButton.innerText = "Executar programa selecionado";
}

function disableRunButton() {
  runButton.disabled = true;
  runButton.innerText = "Executando programa selecionado...";
}

function run() {
  const programName = programSelector.value;
  if (programName === "") {
    alert('Nenhum programa foi selecionado.')
    return;
  }
  
  wasmWorker.postMessage({ type: "run", message: programName });
  textAreaStd.value = `-> Rodando programa ${programName}. Início: ${(new Date()).toLocaleTimeString()}\n`;
};

function updateWebassemblyOutput(message, out="stdout") {
  if (out === "stdout") {
    textAreaStd.value += `${message}\n`;
  } else {
    textAreaStd.value += `----> (${out}) ${message}\n`;
  }
};

function cancelRun() {
  textAreaStd.value = null;
  window.location.reload();
}

runButton.addEventListener('click', () => run());
cancelRunButton.addEventListener('click', () => cancelRun());
