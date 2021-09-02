// При полной загрузке страницы Начинаем загружать Карточки Артистов
var Artists = [
    // Artist Example
    // {
    //     Id:           int
    //     Name:         string
    //     Image:        string
    //     Members:      []string
    //     CreationDate: int
    //     FirstAlbum:   time.Time
    //     Concerts:     [
    //         {
    //             Location    string
    //             Coordinates [2]float64 // latitude first element, longtitude second element
    //             Date        []time.Time //First Concert Date
    //         }
    //     ]
    // }
]

var Suggestions = [
    // Suggestions Example
    // {
    //     "Name": "Name",
    //     "Type": "Member"
    // },
]

var Filter = {
    "Search": "",
    "Filter": {
        "IsEnabled": "false",
        "CreationDate": {
            "IsEnabled": "false",
            "FromYear": 0,
            "BeforeYear": 0
        },
        "Cities": {
            "IsEnabled": "false",
            "Locations": [],
        },
        "FirstAlbumDate": {
            "IsEnabled": "false",
            "FromDate": "",
            "BeforeDate": ""
        },
        "CountMembers": {
            "IsEnabled": "false",
            "From": 0,
            "Before": 0
        }
    }
}

const HostAPI = 'http://localhost:8080/api/'
const HostAPIArtists = HostAPI + 'artists/'
const HostAPIFilter = HostAPI + 'filter/'
const HostAPISuggestions = HostAPI + 'suggestions/'
const CurrentYear = (new Date()).getFullYear();

// Form Objects
const btn_SubmitFilter = document.getElementById('btn_SubmitFilter')
const tb_SearchBar = document.getElementById('tb_SearchBar')
const dl_SearchBar = document.getElementById('dl_SearchBar')

//ShowHideBlocks
const cb_Cities = document.getElementById('cb_Cities')
const b_Cities = document.getElementById('b_Cities')

const cb_CreationDate = document.getElementById('cb_CreationDate')
const b_CreationDate = document.getElementById('b_CreationDate')

const cb_FirstAlbumDate = document.getElementById('cb_FirstAlbumDate')
const b_FirstAlbumDate = document.getElementById('b_FirstAlbumDate')

const cb_CountMembers = document.getElementById('cb_CountMembers')
const b_CountMembers = document.getElementById('b_CountMembers')


window.onload = () => {
    InitHanlers()
    LoadArtists()
    LoadSuggestions()
}

function InitHanlers() {
    btn_SubmitFilter.onclick = LoadArtistsByFilter
    cb_Cities.onchange = cb_CitiesOnChange
    cb_CreationDate.onchange = cb_CreationDateOnChange
    cb_FirstAlbumDate.onchange = cb_FirstAlbumDateOnChange
    cb_CountMembers.onchange = cb_CountMembersOnChange
}

// return Header for correct cors load
function GetCORSHeaders() {
    let headers = new Headers();
    headers.append('Content-Type', 'application/json');
    headers.append('Accept', 'application/json');
    headers.append('Origin', 'http://localhost:3000');
    return headers
}

function ShowErrorOnArtsistBlock(Message, func) {
    var container = document.getElementById("Artists");
    container.innerHTML = `
    <div class="container_error">
        <div class="error_card">
            <div class="w100 mgb">
                <label class="label">Error:</label>
                <p>${Message}</p>
            </div>
            <btn class="btn btn--submit" onclick="${func}()">Try Again</btn>
        </div>
    </div>
    `
}
// Load All Artists
async function LoadArtists() {
    try {
        //Artists
        const response = await fetch(HostAPIArtists, {
            mode: 'cors',
            method: 'GET',
            headers: GetCORSHeaders()
        });
        if (response.status != 200) {
            ShowErrorOnArtsistBlock(`Server error code: ${response.status}\n${response.statusText}`, "LoadArtists")
            return
        }
        // Storing data in form of JSON
        Artists = await response.json()
        await RefreshArtistsOnPage()
    } catch (error) {
        ShowErrorOnArtsistBlock(`The server is not responding`, "LoadArtists")
    }
}
// Обновляет карточки Артистов(Берет данные с Artists)
async function RefreshArtistsOnPage() {
    var container = document.getElementById("Artists");
    container.innerHTML = ""
    Artists.forEach(function (artist, index) {
        // Adding Cards
        var newNode = document.createElement(`div`);
        newNode.className = 'card'
        newNode.innerHTML = `
        <div class="wrap">
            <div class="card-wrap">
                <a><img class="card_img" src="${artist['Image']}" /></a>
            </div>
            <div class="loop-action">
                <a onclick="ShowArtistWindow(${index})">
                    More Details
                </a>
            </div>
        </div>
        <div class="card-info">
            <h3 class="card-title">${artist['Name']}</h3 >
        </div >
        `;
        container.appendChild(newNode);
    });
}
// Load Artists by filter
async function LoadArtistsByFilter() {
    if (!ResetFilterValues()) {
        return
    }
    try {
        const response = await fetch(HostAPIFilter, {
            mode: 'cors',
            method: "POST",
            body: JSON.stringify(Filter),
            headers: GetCORSHeaders(),
        })
        if (response.status != 200) {
            ShowErrorOnArtsistBlock(`Server error code: ${response.status}\n${response.statusText}`, "LoadArtistsByFilter")
            return
        }
        // Storing data in form of JSON
        Artists = await response.json()
        await RefreshArtistsOnPage()
    } catch (error) {
        ShowErrorOnArtsistBlock(`The server is not responding`, "LoadArtistsByFilter")
    }
}
function FilterSetCities() {
    Filter.Filter.Cities.IsEnabled = cb_Cities.checked
    if (cb_Cities.checked) {
        Filter.Filter.Cities.Locations = []
        let allCityCheckBox = document.querySelectorAll('.city')
        allCityCheckBox.forEach((checkbox) => {
            if (checkbox.checked) {
                Filter.Filter.Cities.Locations.push(checkbox.value)
            }
        })
    } else {
        Filter.Filter.Cities.Locations = []
    }
}
function FilterSetCreationDate() {
    Filter.Filter.CreationDate.IsEnabled = cb_CreationDate.checked
    if (cb_CreationDate.checked) {
        let from = document.getElementById('tb_CreationDateFrom').value
        let before = document.getElementById('tb_CreationDateBefore').value
        let fromNum = parseInt(from)
        let beforeNum = parseInt(before)
        if (Number.isNaN(fromNum) || fromNum < 0 || Number.isNaN(beforeNum) || beforeNum < 0) {
            b_CreationDate.style.border = "1px dashed red"
            alert("Incorrect values on Creation Date")
            return false
        }
        b_CreationDate.style.border = "none"

        Filter.Filter.CreationDate.FromYear = fromNum
        Filter.Filter.CreationDate.BeforeYear = beforeNum
    } else {
        Filter.Filter.CreationDate.FromYear = 0
        Filter.Filter.CreationDate.BeforeYear = 0
    }
    return true
}
function FilterSetFirstAlbum() {
    Filter.Filter.FirstAlbumDate.IsEnabled = cb_FirstAlbumDate.checked
    if (cb_FirstAlbumDate.checked) {
        let from = document.getElementById('tb_FirstAlbumDateFrom').value
        let before = document.getElementById('tb_FirstAlbumDateBefore').value
        let date_regex = /^(\d{4})-(\d{2})-(\d{2})$/;

        if (!date_regex.test(from) || !date_regex.test(before)) {
            b_FirstAlbumDate.style.border = "1px dashed red"
            alert('Incorrect First Album Date Value')
            return false
        }
        b_FirstAlbumDate.style.border = "none"

        let fromDate = new Date(from)
        let beforeDate = new Date(before)
        if (!isValidDate(fromDate) || !isValidDate(beforeDate)) {
            alert('Incorrect First Album Date Value')
            return false
        }

        Filter.Filter.FirstAlbumDate.FromDate = fromDate
        Filter.Filter.FirstAlbumDate.BeforeDate = beforeDate
    } else {
        Filter.Filter.FirstAlbumDate.FromDate = ""
        Filter.Filter.FirstAlbumDate.BeforeDate = ""
    }
    return true
}
function FilterSetCountMembers() {
    Filter.Filter.CountMembers.IsEnabled = cb_CountMembers.checked
    if (cb_CountMembers.checked) {
        let from = document.getElementById('tb_CountMembersFrom').value
        let before = document.getElementById('tb_CountMembersBefore').value

        let fromNum = parseInt(from)
        let beforeNum = parseInt(before)
        if (Number.isNaN(fromNum) || fromNum < 0 || Number.isNaN(beforeNum) || beforeNum < 0) {
            b_CountMembers.style.border = "1px dashed red"
            alert("Incorrect values on Count Members")
            return false
        }
        b_CountMembers.style.border = "none"

        Filter.Filter.CountMembers.From = fromNum
        Filter.Filter.CountMembers.Before = beforeNum
    } else {
        Filter.Filter.CountMembers.From = 0
        Filter.Filter.CountMembers.Before = 0
    }
    return true
}
function isValidDate(d) {
    return d instanceof Date && !isNaN(d);
}

function ResetFilterValues() {
    Filter.Search = tb_SearchBar.value

    Filter.Filter.IsEnabled = (
        cb_Cities.checked ||
        cb_CreationDate.checked ||
        cb_FirstAlbumDate.checked ||
        cb_CountMembers.checked
    );
    FilterSetCities()
    return (
        FilterSetCreationDate() &&
        FilterSetFirstAlbum() &&
        FilterSetCountMembers()
    )
}

// Search Bar And Cities Data
async function LoadSuggestions() {
    try {
        //Suggestions
        const response = await fetch(HostAPISuggestions, {
            mode: 'cors',
            method: 'GET',
            headers: GetCORSHeaders()
        });

        // Storing data in form of JSON
        Suggestions = await response.json();
        Suggestions.sort(function (a, b) {
            return (a.Name).localeCompare(b.Name);
        })
        await RefreshSearchBar()
        await RefreshCityCheckBoxs()
        let filterErrorBlock = document.getElementById('b_FilterError')
        filterErrorBlock.style.display = "none"
        filterErrorBlock.innerHTML = ``
    } catch (error) {
        // alert('Error with getting Suggestions data')
        let filterErrorBlock = document.getElementById('b_FilterError')
        filterErrorBlock.style.display = "block"
        filterErrorBlock.innerHTML = `
            <label class="label">Error with load filter settings</label>
            <btn class="btn btn--submit" style="width: 100%;" onclick="LoadSuggestions()">Try Again</btn>
        `
    }
}
async function RefreshSearchBar() {
    dl_SearchBar.innerHTML = ""
    Suggestions.forEach((suggestion) => {
        let option = document.createElement('option')
        option.value = suggestion.Name
        option.label = `${suggestion.Name} - ${suggestion.Type}`
        dl_SearchBar.appendChild(option);
    })
}
async function RefreshCityCheckBoxs() {
    b_Cities.innerHTML = ""
    Suggestions.forEach((suggestion) => {
        if (suggestion.Type == "Concert Location") {
            let label = document.createElement('label')
            label.innerHTML = `<label><input type="checkbox" value="${suggestion.Name}" class="city"> ${suggestion.Name}</label><br>`
            b_Cities.appendChild(label);
        }
    })
}

// SHOW\HIDE EVENTS
function cb_CitiesOnChange(event) {
    if (event.target.checked) {
        b_Cities.style.maxHeight = `240px`
        b_Cities.style.padding = `5px`
        b_Cities.style.margin = `0 0 10px 0`

        Filter.Filter.Cities.IsEnabled = true
    } else {
        b_Cities.style.maxHeight = `0px`
        b_Cities.style.padding = `0px`
        b_Cities.style.margin = `0px`

        Filter.Filter.Cities.IsEnabled = false
    }
}
function cb_CreationDateOnChange(event) {
    if (event.target.checked) {
        b_CreationDate.style.maxHeight = `240px`
        b_CreationDate.style.margin = `0 0 10px 0`

        Filter.Filter.CreationDate.IsEnabled = true
    } else {
        b_CreationDate.style.maxHeight = `0px`
        b_CreationDate.style.margin = `0px`
        b_CreationDate.style.border = "none"

        Filter.Filter.CreationDate.IsEnabled = false
    }
}
function cb_CountMembersOnChange(event) {
    if (event.target.checked) {
        b_CountMembers.style.maxHeight = `240px`
        b_CountMembers.style.margin = `0 0 10px 0`

        Filter.Filter.CreationDate.IsEnabled = true
    } else {
        b_CountMembers.style.maxHeight = `0px`
        b_CountMembers.style.margin = `0px`
        b_CountMembers.style.border = "none"

        Filter.Filter.CountMembers.IsEnabled = false
    }
}
function cb_FirstAlbumDateOnChange(event) {
    if (event.target.checked) {
        b_FirstAlbumDate.style.maxHeight = `240px`
        b_FirstAlbumDate.style.margin = `0 0 10px 0`

        Filter.Filter.CreationDate.IsEnabled = true
    } else {
        b_FirstAlbumDate.style.maxHeight = `0px`
        b_FirstAlbumDate.style.margin = `0px`
        b_FirstAlbumDate.style.border = "none"

        Filter.Filter.FirstAlbumDate.IsEnabled = false
    }
}
// POPUP
function ShowArtistWindow(index) {
    SetArtistToWindow(index)

    // Show popup
    document.getElementById('page').style.overflow = 'hidden'
    document.getElementById('window').style.top = `${window.scrollY}px`;
    document.getElementById('window').style.height = '100vh'
    document.getElementById('window').style.padding = '10px'
}
function HideArtistWindow() {
    document.getElementById('page').style.overflow = 'auto'
    document.getElementById('window').style.height = '0vh'
    document.getElementById('window').style.padding = '0px'
}
// Refresh Window Information
function SetArtistToWindow(index) {
    let artist = Artists[index]

    document.getElementById('ArtistName').innerHTML = artist['Name']
    document.getElementById('ArtistImage').src = artist['Image']
    document.getElementById('ArtistCreationDate').innerHTML = `${artist['CreationDate']}`
    document.getElementById('ArtistFirstAlbum').innerHTML = (new Date(artist['FirstAlbum']).toDateString())
    //Set Members
    let membersBlock = document.getElementById('ArtistMembers')
    membersBlock.innerHTML = `<label class="label">Members</label>`
    artist.Members.forEach((member) => {
        let newNode = document.createElement(`p`);
        newNode.innerHTML = `${member}`
        membersBlock.appendChild(newNode)
    });

    //Set Locations
    Locations = []
    var container = document.getElementById("ArtistConcerts");
    container.innerHTML = ""
    artist.Concerts.forEach((concert) => {
        // Set Dates
        let newBlock = document.createElement(`div`);
        newBlock.className = 'group-info'
        newBlock.innerHTML = `<label class="sublabel">${concert.Location}</label>`
        concert.Date.forEach((concertDate) => {
            let formattedDate = (new Date(concertDate).toDateString())
            let newNode = document.createElement(`p`);
            newNode.innerHTML = `${formattedDate}`
            newBlock.appendChild(newNode)
        });
        container.appendChild(newBlock);

        // Set Locations For Google Map
        Locations.push({
            lat: concert['Coordinates'][0],
            lng: concert['Coordinates'][1],
        })
    });
    // Refresh Map
    myMap()
}