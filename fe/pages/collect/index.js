let {API} = require('../../utils/api')
Page({
  data: {
    folderList: [],
    paperList: [],
    folderName: '',
    selected: 0,
    to_add_folder: false,
    input: ''
  },
  change: function(e) {
    this.setData({
      folderName: this.data.folderList[e.currentTarget.dataset.index].name,
      selected: e.currentTarget.dataset.index,
      paperList: this.data.folderList[e.currentTarget.dataset.index].content
    })
  },
  jump_collect_edit(e) {
    if(!this.data.edit) {
      wx.navigateTo({
      url: '/pages/collect_edit/index',
    })
    }
  },
  getData() {
    let that = this
    API.getFav().then(res => {
      console.log(res);
      let list = res.data.data;
      let selected = this.data.selected;
      let defaultFolderIndex = list.findIndex(item => {
        return item.name == '默认收藏'
      });
      let defaultFolder = list.splice(defaultFolderIndex,1)[0];
      list.unshift(defaultFolder)
      list.forEach(item => {
        item.content.forEach(item2 => {
          item2['choosen'] = false;
          item2.authors = item2.authors.slice(0,4);
        })
      })
      that.setData({
        folderList: list,
        paperList: list[selected].content,
        folderName: list[selected].name,
        edit: false,
        count: 0
      })
    })
  },
  onShow: function(e) {
    this.getData()
  },
  
  // to_collect
  cancel: function() {
    let paperList = this.data.paperList;
    paperList.forEach(item => {
      item.choosen = false
    });
    this.setData({
      paperList,
      edit: false,
      count: 0
    })
  },
  all: function() {
    let paperList = this.data.paperList;
    let count = 0;
    paperList.forEach(item => {
      item.choosen = true
      count++
    });
    this.setData({
      paperList,
      edit: true,
      count
    })
  },
  dotTap(e) {
    let id = e.currentTarget.dataset.id;
    wx.navigateTo({
      url: '/pages/collect_paper_edit/index?name='+this.data.folderName+'&id='+id,
    })
  },

  // add_folder
  add_folder(e) {
    this.setData({
      to_add_folder: true
    })
  },
  cancel_add_folder(e) {
    this.setData({
      to_add_folder: false
    })
  },
  bindInput(e) {
    this.setData({
      input: e.detail.value
    })
  },
  new_folder() {
    API.addFavFolder(this.data.input).then(res => {
      console.log(res);
      this.setData({
        input: '',
        to_add_folder: false
      })
      wx.reLaunch({
        url: '/pages/collect/index',
      })
    })
  },

  // search
  jump_collect_search(e) {
    wx.navigateTo({
      url: '/pages/collect_search',
    })
  }
})