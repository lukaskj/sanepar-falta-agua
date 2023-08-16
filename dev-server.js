const http = require("http");

const host = 'localhost';
const port = 8000;

const defaultResponse = {
  "Mensagem":"",
  "PrevisaoData":"15\/08\/2023",
  "PrevisaoHora":"09:00",
  "NormalizacaoData":"15\/08\/2023",
  "NormalizacaoHora":"15:00",
}


const requestListener = function (req, res) {
    res.writeHead(200);
    res.end(JSON.stringify(defaultResponse));
};

const server = http.createServer(requestListener);
server.listen(port, host, () => {
    console.log(`Server is running on http://${host}:${port}`);
});