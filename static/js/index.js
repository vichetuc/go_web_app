$(document).ready(function() {
    $.ajax({
        url: "/greetings"
    }).then(function(data) {
        var greetingsHtml = data.map(greeetingTemplate);
        $(document.body).append("<ul>" + greetingsHtml + "</ul>");
    });
});

function greeetingTemplate(greeting) {
    return "<li>" + greeting.author + ":" + greeting.content +"</li>";
}