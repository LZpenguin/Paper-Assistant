
function get_doc(id,token,callback) {
  wx.request({
    url: 'https://api.hust.online/doc-aid-test/fav',
    success: res => {
      let folderList = res.data.data;
      let folderName = folderList.find(item => {
        if(item.content.find(item2 => item2.id == id))
          return true
      }).name
      callback(folderName)
    },
    header: {
      'Authorization': 'Bearer '+ token
    }
  })
}

function fav_add(doc,id,token,callback) {
  wx.request({
    url: 'https://api.hust.online/doc-aid-test/fav/'+doc,
    method: 'post',
    data: [id],
    success: res => {
      if(callback) {
        callback(res)
      }
    },
    header: {
      'Authorization': 'Bearer '+ token
    }
  })
}

function fav_del(id,token,callback) {
  get_doc(id,token,doc => {
    wx.request({
      url: 'https://api.hust.online/doc-aid-test/fav/'+doc,
      method: 'delete',
      data: [id],
      success: res => {
        if(callback) {
          callback(res)
        }
      },
      header: {
        'Authorization': 'Bearer '+ token
      }
    })
  })
}

module.exports = {get_doc,fav_add,fav_del}