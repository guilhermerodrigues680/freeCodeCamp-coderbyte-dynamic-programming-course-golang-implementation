// importScripts('wasm_exec_go.js'); // glue js file for go wasm
importScripts('wasm_exec_tinygo.js'); // glue js file for tinygo wasm

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
  webassemblyCompletedSuccessfully: 'webassembly_completed_successfully',
  webassemblyTerminatedWithError: 'webassembly_terminated_with_error',
  genericError: 'generic_error',
};
const go = new Go();

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

async function run(programName) {
  const programWasm = `${programName}/main.wasm`;
  let result;
  try {
    // TODO: Implementar rotina de cache
    const programPath = `../${programWasm}`;
    result = await WebAssembly.instantiateStreaming(fetch(programPath), go.importObject);
    postMessageOk("webassembly started", MESSAGE_TYPES.webassemblyStarted);
  } catch (err) {
    console.error(err);
    return;
  }

  postMessageOk(`running webassembly ${programWasm}`, MESSAGE_TYPES.runningWebassembly);

  try {
    const r = await go.run(result.instance);
    console.log(r);
  } catch (error) {
    const errStr = `${error.name}: ${error.message}`;
    postMessageError(`webassembly ${programWasm} terminated with error: ${errStr}`, MESSAGE_TYPES.webassemblyTerminatedWithError);
    postMessageError(errStr, MESSAGE_TYPES.writeToStdout);
    // console.debug('go.run error',error);
    return;
  }

  // TODO: reset instance
  // inst = await WebAssembly.instantiate(mod, go.importObject); // reset instance
  postMessageOk(`webassembly ${programWasm} completed successfully`, MESSAGE_TYPES.webassemblyCompletedSuccessfully);
}

onmessage = async function(e) {
  // console.debug('Worker: Message received from main script');
  const message = e.data;
  if (message?.type) {
    // console.debug(`${message.type} message`);
    switch (message.type) {
    case 'init':
      initWasm();
      break;
    case 'run':
      run(message.message);
      break;
    default:
      console.debug(`Mensagem ${message.type} desconhecida.`);
      break;
    }
  } else {
    console.debug(`Mensagem ${message} não possui tipo.`);
  }
}
