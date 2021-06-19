importScripts('wasm_exec_tinygo.js');

if (!WebAssembly.instantiateStreaming) { // polyfill
  WebAssembly.instantiateStreaming = async (resp, importObject) => {
    const source = await (await resp).arrayBuffer();
    return await WebAssembly.instantiate(source, importObject);
  };
}

const MESSAGE_TYPES = {
  writeToStdout: 'write_to_stdout',
  webassemblyStarted: 'webassembly_started',
  runningWebassembly: 'running_webassembly',
  webassemblyExecuted: 'webassembly_executed',
  genericError: 'generic_error',
};
const go = new Go();
let mod, inst;

// Função Global de callback do `js/wasm_exec_tinygo.js` ou `js/wasm_exec_go.js`
self.GO_WEBASSEMBLY_STD_CALLBACK = function (line) {
  postMessageOk(line, MESSAGE_TYPES.writeToStdout);
};

function postMessageOk(message, type) {
  postMessage({ type, message });
}

function postMessageError(error, type) {
  postMessage({ type, error });
}

function initWasm() {
  WebAssembly.instantiateStreaming(fetch("../01-fib/main.wasm"), go.importObject).then((result) => {
    mod = result.module;
    inst = result.instance;
    postMessageOk("webassembly started", MESSAGE_TYPES.webassemblyStarted);
  }).catch((err) => {
    console.error(err);
  });
}

async function run() {
  postMessageOk("running webassembly", MESSAGE_TYPES.runningWebassembly);
  await go.run(inst);
  inst = await WebAssembly.instantiate(mod, go.importObject); // reset instance
  postMessageOk("webassembly executed", MESSAGE_TYPES.webassemblyExecuted);
}

onmessage = async function(e) {
  console.debug('Worker: Message received from main script');
  const message = e.data;
  if (message?.type) {
    console.debug(`${message.type} message`);
    switch (message.type) {
    case 'init':
      initWasm();
      break;
    case 'run':
      run();
      break;
  
    default:
      console.debug(`Mensagem ${message.type} desconhecida.`);
      break;
    }
  } else {
    console.debug(`Mensagem ${message} não possui tipo.`);
  }
}
