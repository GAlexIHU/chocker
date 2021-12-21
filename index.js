const { spawn } = require('child_process');
const { mkdirSync, existsSync } = require('fs');
const { join } = require('path');

async function main() {
  if (!existsSync(join(__dirname, 'ubuntu'))) {
    mkdirSync(join(__dirname, 'ubuntu'));
  }

  const proc = spawn(
    'sh',
    [
      '-c',
      'archivemount ubuntu.tar ./ubuntu && go build && sudo ./chocker run /bin/bash && umount ./ubuntu || umount ./ubuntu',
    ],
    { cwd: __dirname }
  );
  return new Promise((res, rej) => {
    proc.stdout.on('data', (data) => {
      process.stdout.write(data);
    });

    proc.stderr.on('data', (data) => {
      process.stderr.write(data);
    });

    process.stdin.on('data', (data) => {
      proc.stdin.write(data);
    });

    proc.on('disconnect', (code, signal) => {
      console.log('disconnected');
      res([code, signal]);
    });

    proc.on('exit', (code, signal) => {
      console.log('exiting');
      res([code, signal]);
    });

    proc.on('error', (err) => {
      rej(err);
    });
  });
}

function catcher(e) {
  console.error(e);
  process.exit(1);
}

if (require.main === module) {
  main()
    .then((v) => {
      console.log(v) && process.exit(0);
    })
    .catch(catcher);
}
