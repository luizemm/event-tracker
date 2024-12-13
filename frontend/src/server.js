const express = require("express");
const fs = require("fs");

const app = express();
const port = process.env.PORT || 3000;

app.set("view engine", "ejs");
app.set("views", "./views");

app.get("/", (_req, res) => {
  res.render("example-page");
});

app.get("/monitor", (_req, res) => {
  res.render("app-events-monitor", {
    WS_RECEIVER_URL: process.env.WS_RECEIVER_URL || "",
  });
});

app.get("/ws-sender.js", (_req, res) => {
  const senderJs = fs.readFileSync("./static/js/ws-sender.js", "utf-8");

  res.type("javascript");
  res.send(
    senderJs.replace("%%WS_SENDER_URL%%", process.env.WS_SENDER_URL || "")
  );
});

app.listen(port, () => {
  console.log(`App listening on port ${port}`);
});
