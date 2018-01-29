
window.socket = new WebSocket("ws://" + location.host +"/ws");

function sendMessage(msg)
{
	socket.send(msg)
}

function setUpSocket(onmessage)
{
	socket.onopen = function() 
	{
	  console.log("Connected");
	}
	socket.onclose = function(event) 
	{
	  if (event.wasClean) {
	    console.log('Connection closed');
	  } else {
	    console.log('Error: Connection reset'); // например, "убит" процесс сервера
	  }
	  console.log('Code: ' + event.code + ' reason: ' + event.reason);
	};

	socket.onmessage = function(event) 
	{
	  console.log("Получены данные " + event.data);
	};

	socket.onmessage = onmessage;

	socket.onerror = function(error) 
	{
	  console.log("Ошибка " + error.message);
};
}
function DisplayMessage (msg)
{
	var container = document.getElementById("container");
	var div = document.createElement("div");
	div.className = 'message';
	var textNode = document.createTextNode(msg);

	div.appendChild(textNode);
	container.appendChild(div);

}