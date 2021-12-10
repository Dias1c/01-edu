// TAG EDITOR
class TagEditor {
    // ! HTML Elements
    B_TagEditor; // new HTMLDivElement();
    Tb_Input; // new HTMLInputElement();
    Tb_Tags; // new HTMLInputElement();

    // ! SETTIGNS Default
    SB_sName = '#b_TagEditor';      // visible TagEditor block SelectorName
    STb_sName = '#tb_TagEditor';    // input SelectorName
    SHasDoubles = false;            // AddDoubles?
    SToLower = true;                // ToLowercase Tag?
    SMaxTags = 0;                   // Max Tags Count, 0 == Unlimited
    STags = [];                     // Writed Tags

    // Inits TagEditor By Settings
    constructor(settings) {
        // Set Settings
        this.init_Settings(settings);
        // Init Inputs and Blocks on html
        this.B_TagEditor = document.querySelector(this.SB_sName);
        this.Tb_Tags = document.querySelector(this.STb_sName);
        this.init_Tb_Input();
        // Set Tags On input and STags
        this.RefreshTags();
    }
    // Set Default Settings by Settings
    init_Settings(settings) {
        this.SB_sName = (settings.BlockSelectorName) ? settings.BlockSelectorName : this.SB_sName;
        this.STb_sName = (settings.TextBlockSelectorName) ? settings.TextBlockSelectorName : this.STb_sName;
        this.SHasDoubles = (settings.HasDoubles) ? settings.HasDoubles : this.SHasDoubles;
        this.SToLower = (settings.ToLower) ? settings.ToLower : this.SToLower;
        this.SMaxTags = (settings.MaxTags) ? settings.MaxTags : this.SMaxTags;
        this.STags = (settings.Tags) ? settings.Tags : this.STags;
    }
    init_Tb_Input() {
        if (this.B_TagEditor.querySelector('input') == null) {
            this.B_TagEditor.appendChild(document.createElement('input'));
        }
        this.Tb_Input = this.B_TagEditor.querySelector('input');

        // Add Events
        this.Tb_Input.addEventListener('keyup', (e) => {
            if (e.key === ' ') {
                this.Tb_Input.value = '';
            }
        });
        this.Tb_Input.addEventListener('keydown', (e) => {
            if (e.key === ' ' || e.key === 'Enter') {
                let tagname = this.Tb_Input.value;
                console.log(`TagName: "${tagname}"`)
                this.AddTag(tagname);
                this.Tb_Input.value = '';
            } else if (e.key === 'Backspace' && this.Tb_Input.value == "") {
                console.log('Remove Tag');
                let tagname = this.RemoveLastTag();
                this.Tb_Input.value = tagname + " ";
            }
        });

        // Add Tags
        this.Tb_Input.value.split().forEach((tagname, index) => {
            if (!this.IsValidName(tagname)) {
                return
            }
            if (this.SToLower) {
                tagname = tagname.toLowerCase();
            }
            this.STags.push(tagname);
        });
        this.Tb_Input.value = '';
    }

    // ! Public Methods
    IsValidName(name) {
        if (this.SToLower) {
            name = name.toLowerCase();
        }
        if (name == "") {
            return false
        } else if (this.SMaxTags != 0 && this.STags.length >= this.SMaxTags) {
            console.log("SMaxTags", name);
            return false
        } else if (!this.SHasDoubles && this.STags.includes(name)) {
            console.log("SHasDoubles", name);
            return false
        }
        return true
    }
    // Add Tag to TagEditor
    AddTag(name) {
        if (!this.IsValidName(name)) {
            return
        }
        if (this.SToLower) {
            name = name.toLowerCase();
        }
        this.STags.push(name);
        console.info(this.STags);
        this.RefreshTags();
    }
    // Removes last tag from TagEditor
    RemoveLastTag() {
        let name = this.STags.pop();
        this.RefreshTags();
        return (name != null) ? name : "";
    }
    // Removes tag by name from TagEditor
    RemoveTag(name) {
        this.STags = this.STags.filter(function (value, index, arr) {
            return value != name;
        });
        this.RefreshTags();
        return (name != null) ? name : "";
    }
    // Sets Tags to TagEditor by STags
    RefreshTags() {
        while (this.B_TagEditor.firstChild != this.Tb_Input) {
            this.B_TagEditor.removeChild(this.B_TagEditor.firstChild);
        }
        this.STags.slice().reverse().forEach((tagname, index) => {
            const newtag = TagEditor.CreateTag(tagname);
            this.B_TagEditor.prepend(newtag);
        });
        this.Tb_Tags.value = this.STags.join(' ');
    }

    // ! Static Methods

    // Returns HtmlElement Tag
    static CreateTag(name) {
        const btn_tag = document.createElement('div');
        btn_tag.setAttribute('class', 'btn-tag');
        const tag_name = document.createElement('span');
        tag_name.innerHTML = name
        const btn_remove = document.createElement('span');
        btn_remove.setAttribute('class', 'remove');
        btn_remove.innerHTML = '×';
        btn_remove.addEventListener('click', () => {
            btn_tag.remove();
        });
        // Construct tag
        btn_tag.appendChild(tag_name);
        btn_tag.appendChild(btn_remove);
        return btn_tag
    }
}
