const socket = new WebSocket("%%WS_SENDER_URL%%");

function sendMessage(event, data) {
  if (socket.readyState === socket.OPEN) {
    socket.send(
      JSON.stringify({
        event_type: event.type,
        data: JSON.stringify(data),
        timestamp: new Date().toISOString(),
      })
    );
  }
}

function sendMessageWithType(eventType, data) {
  if (socket.readyState === socket.OPEN) {
    socket.send(
      JSON.stringify({
        event_type: eventType,
        data: JSON.stringify(data),
        timestamp: new Date().toISOString(),
      })
    );
  }
}

socket.addEventListener("open", (event) => {
  console.log("websocket connection opened");

  const parameters = Object.fromEntries(
    new URLSearchParams(window.location.search)
  );

  sendMessage(event, { parameters });
  sendMessageWithType("page_view", {
    size: {
      width: window.screen.width,
      height: window.screen.height,
    },
  });
});

document.addEventListener("click", (event) => {
  sendMessage(event, {
    element: event.target.outerHTML,
  });
});

const lastScrollEventPosition = {
  x: 0,
  y: 0,
};

document.addEventListener("scroll", (event) => {
  if (
    Math.abs(lastScrollEventPosition.x - window.scrollX) > 100 ||
    Math.abs(lastScrollEventPosition.y - window.scrollY) > 100
  ) {
    lastScrollEventPosition.x = window.scrollX;
    lastScrollEventPosition.y = window.scrollY;

    sendMessage(event, {
      position: {
        x: window.scrollX,
        y: window.scrollY,
      },
    });
  }
});

window.addEventListener("resize", (event) => {
  sendMessage(event, {
    size: {
      width: window.innerWidth,
      height: window.innerHeight,
    },
  });
});
