const conversation = $("#list"); // $ means jquery, it won't get the item without the #
const userInput = $("#userInput");
 console.log(conversation);


userInput.keypress(function(event){
    if(event.keyCode !== 13){ // 13 is the keycode for Enter
        return; // do nothing unless the key is enter
    }
    event.preventDefault(); // prevents the form default behaviour which would refresh the page.
    const text = userInput.val();

    console.log(text);
    userInput.val(""); // set it to nothing, .val() is like a getter, .val(" ") is like a setter

    // trim removes all spaces from either side,
    // if there's no text left, the user doesn't have a question.
    if(text.trim() == ""){
    //if(!text.trim()){ 
        return;
    }

    // a query parameter user-input is expected
    queryParameters = {
        "userInput" : text
    }

    conversation.append("<li class=\"list-group\">" + text + "<li class=\"list-group\">");

    // sample url generated.s 
    //http://localhost:8080/ask?user-input=hello%20world
    $.get("/ask", queryParameters).done(function(resp){
        // this code will execute when the request gets a response.
        setTimeout(function(){ // wait 1 second then add element.
            conversation.append("<li class=\"list-group\">" + resp + "<li class=\"list-group\">");
        }, 1000);
        
    }).fail(function(){ // this will run whenever anything goes wrong.
        conversation.append("<li class=\"list-group\">the doctor is out, sorry.</li class=\"list-group\">");
    });
});
