
Page({
  data: {

  },
  onShow: function (options) {
    wx.getStorage({
      key: 'shared',
      success: res => {
        this.setData({
          sharedList: res.data,
          result: res.data
        })
      }
    })
  },
  getSearch(e) {
    let value = e.detail.value.trim();
    let reg = new RegExp(value);
    let list = this.data.sharedList;
    let result = [];
    list.forEach(item => {
      if(item.title.match(reg)) {
        result.push(item)
      }
    })
    if(value == '') {
      result = list;
    }
    this.setData({
      result
    })
  }
})