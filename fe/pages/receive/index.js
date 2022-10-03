// pages/receive/index.js
Page({
  data: {
    paperList: [],
    count: 0,
    edit: false,
    to_move: false,
    to_add_folder: false
  },
  getData(cb) {
    var that = this;
    wx.getStorage({
      key: 'token',
      success: res => {
        this.setData({
          'token': res.data
        })
        wx.request({
          url: 'https://api.hust.online/doc-aid-test/fav',
          success: function(rel) {
            let folderList = rel.data.data;
            that.setData({
              folderList
            })
          },
          failed: function (err) {
            setTimeout(() => {
              this.getData()
            }, 200);
          },
          header: {
            'Authorization': 'Bearer '+ res.data
          }
        })
        if(cb) {
          cb()
        }
      }
    })
  },
  onLoad(e) {
    let idList = e.ids.split('+');
    this.getData(() => {
      idList.forEach(item => {
       wx.request({
         url: 'https://api.hust.online/doc-aid-test/document/'+item,
         success: (rel) => {
           let paperList = this.data.paperList;
           let paper = rel.data.data;
           paper.authors = paper.authors.slice(0,4);
           paper.suggest.forEach(item => {
             item.authors = item.authors.slice(0,4);
           })
           paperList.push(paper);
           console.log(paperList);
           this.setData({
             paperList,
             nickName: e.nickName
           })
         },
         header: {
           'Authorization': 'Bearer '+ this.data.token
         }
       })
      })
    })
  },
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
  dotTap: function(e) {
    let edit = false;
    let count = this.data.count;
    let paperList = this.data.paperList;
    let item = paperList.find(item => item.id == e.currentTarget.dataset.id);
    if(item.choosen) {
      item.choosen = false;
      count--;
    }else {
      item.choosen = true;
      count++;
    }
    if(paperList.find(item => item.choosen == true)) {
      edit = true;
    }
    this.setData({
      paperList,
      edit,
      count
    })
  },
  cancel_move() {
    this.setData({
      to_move: false
    })
  },
  move_to(e) {
    let paperList = this.data.paperList;
    let delList = [];
    paperList.forEach(item => {
      if(item.choosen) {
        delList.push(item.id)
      }
    });
    wx.request({
      url: 'https://api.hust.online/doc-aid-test/fav/'+e.currentTarget.dataset.name,
      method: 'post',
      data: delList,
      success: (rel) => {
        if(this.data.jump_back) {
          wx.navigateBack()
        }else {
          wx.reLaunch({
            url: '/pages/collect/index',
          })
        }
      },
      header: {
        'Authorization': 'Bearer '+ this.data.token
      }
    })
  },
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
    wx.request({
      url: 'https://api.hust.online/doc-aid-test/folder/'+this.data.input,
      method: 'post',
      success: (rel) => {
        wx.request({
          url: 'https://api.hust.online/doc-aid-test/fav',
          success: (rel) =>{
            this.setData({
              folderList: rel.data.data
            })
          },
          header: {
            'Authorization': 'Bearer '+ this.data.token
          }
        })
        this.setData({
          input: '',
          to_add_folder: false
        })
      },
      header: {
        'Authorization': 'Bearer '+ this.data.token
      }
    })
  },
  collect(e) {
    this.setData({
      to_move: true
    })
  }
})