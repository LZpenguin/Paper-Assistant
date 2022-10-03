// pages/collect_seach/index.js
Page({
  data: {
    folderList: [],
    result: []
  },
  getData() {
    var that = this;
    wx.getStorage({
      key: 'token',
      success: res => {
        this.setData({
          'token': res.data
        })
        wx.request({
          url: 'https://api.hust.online/doc-aid-test/fav',
          success: (rel) => {
            let list = rel.data.data;
            list.forEach(item => {
              item.content.forEach(item2 => {
                item2.authors = item2.authors.slice(0,4)
              })
            })
            that.setData({
              folderList: list
            })
          },
          header: {
            'Authorization': 'Bearer '+ res.data
          }
        })
      }
    })
  },
  onLoad(e) {
    this.getData()
  },
  getSearch(e) {
    let reg = new RegExp(e.detail.value);
    let list = this.data.folderList;
    let result = [];
    list.forEach(item => {
      item.content.forEach(item2 => {
        if(item2.title.match(reg)) {
          result.push(item2)
        }
      })
    })
    this.setData({
      result
    })
  }
})