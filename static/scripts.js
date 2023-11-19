// API

export function getAllAudioFiles() {
    var tracksArray = [];
    $.ajax({
        url: "api/get_audio_files/",
        type: "GET",
        contentType: "application/json",
        success: function(data) {
            data["mp3_files"].forEach(item => { tracksArray.push(item); });
        },
    });
    return tracksArray;
}

