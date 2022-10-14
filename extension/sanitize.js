const htmlFile = process.argv[2];

if (!htmlFile) {
    process.stderr.write("require html file parameter\n");
    process.exit(1);
}

const fs = require('fs');

if (!fs.existsSync(htmlFile)) {
    process.stderr.write("file not exists\n");
    process.exit(1);
}

let html;
try {
    html = fs.readFileSync(htmlFile, 'utf8');
} catch (err) {
    process.stderr.write("cannot read file ", err, "\n");
    process.exit(1);
}

html = html.replace(/"\/static\//g, "\"static/");

try {
    fs.writeFileSync(htmlFile, html);
  } catch (err) {
    process.stderr.write("cannot write to file", err, "\n");
    process.exit(1);
  }
