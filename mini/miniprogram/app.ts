import { HOST_NAME } from "./model/constants"

// app.ts
App<IAppOption>({
  globalData: {},
  onLaunch() {

    // 登录
    wx.login({
      success: data => {
        wx.request(
          {
            url: `${HOST_NAME}/user/login`, method: 'POST',
            data: data,
            success: (res) => {
              console.log(res)
            }
          }
        )
      },
    })
  },
})