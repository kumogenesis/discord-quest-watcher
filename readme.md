# Korchi Tracker - Mô tả cách hoạt đông

<img width="400" height="248" alt="image" src="https://github.com/user-attachments/assets/3771e972-9898-450a-b1c5-32aff759fbc8" />

Nhận thông báo khi có nhiệm vụ Discord mới được phát hành. Lọc theo nhiệm vụ nhận Orb hoặc theo dõi tất cả các loại nhiệm vụ. Ứng dụng Go tối giản, chỉ cần một phụ thuộc duy nhất, đăng nhập đáng tin cậy thông qua mã thông báo người dùng, bỏ qua captcha và giới hạn tốc độ. Hoàn toàn tự lưu trữ và riêng tư. 

> Tại sao?

Tôi không tìm thấy công cụ nào hiện có thể thông báo cho tôi một cách đáng tin cậy khi có nhiệm vụ Orb mới, vì vậy tôi đã tạo ra công cụ nhỏ này trong một buổi tối. 

> Nhưng tại sao người ta lại quan tâm đến những nhiệm vụ đó ngay từ đầu?

Nhiệm vụ cho Orb -> Orb cho những vật phẩm lấp lánh miễn phí. Không, nhưng nghiêm túc mà nói, tôi thích vẻ ngoài của một số vật phẩm trang trí người dùng Discord, nhưng tôi sẽ không bao giờ trả tiền cho nó. À mà nhân tiện, bạn không cần phải *hoàn thành* nhiệm vụ đâu, đoạn mã này sẽ làm điều đó thay bạn: https://gist.github.com/aamiaa/204cd9d42013ded9faf646fae7f89fbb 
Hoặc bạn có thể dùng plugin vencord/equicord/hoặc sử dụng Questor Bot (Có thể search trên mạng để vào server Discord)phổ biến hiện nay.

## Biến môi trường hoạt động

| Biến             | Yêu cầu | Mặc định | Mô tả                                                                 |
|-----------------------|----------|---------|-----------------------------------------------------------------------------|
| `TOKEN`               | Có      | —       | Token Discord User (không khuyến khích)                                                |
| `DISCORD_WEBHOOK_URL` | Có      | —       | webhook URL used for sending notifications                                  |
| `REWARD_FILTER`       | Không       | `Tất cả`   | Bộ lọc phần thưởng: `orbs` (chỉ orbs) hoặc `all` (bao gồm tất cả phần thưởng)       |
| `FETCH_INTERVAL`      | Không       | `30`    | Khoảng thời gian (tính bằng phút) giữa các lần kiểm tra nhiệm vụ (phải là số nguyên dương)      |
| `RUN_ONCE`            | Không       | `Sai` | Nếu `true`, ứng dụng sẽ chạy một lần rồi thoát.                   |
| `WEBHOOK_MESSAGE`     | Có       | —       | Văn bản bổ sung được thêm vào thông báo (ví dụ: thông báo vai trò)                |

### Tin nhắn webhook tùy chỉnh 

Thông báo được gửi dưới dạng nhúng Discord. Bạn có thể thêm tin nhắn văn bản phía trên phần nhúng bằng cách sử dụng biến môi trường `WEBHOOK_MESSAGE`. 

Ví dụ: ping một vai trò

```shell
WEBHOOK_MESSAGE=<@&`1234567890123456789`>
```

## usage

> [!CẢNH BÁO] > Ứng dụng này sử dụng mã thông báo người dùng của bạn, về mặt kỹ thuật, điều này vi phạm Điều khoản dịch vụ của Discord, vì vậy hãy sử dụng với rủi ro của riêng bạn. [Cách lấy mã thông báo người dùng của bạn](https://gist.github.com/MarvNC/e601f3603df22f36ebd3102c501116c6#file-get-discord-token-from-browser-md)

> Hoặc bạn sử dụng JavaScript ở dưới để copy Tok3n Acc0unt Acce55

```shell
javascript:(function(){try{let f=document.createElement('iframe');document.body.appendChild(f);let t=JSON.parse(f.contentWindow.localStorage.token);let ta=document.createElement('textarea');ta.value=t;document.body.appendChild(ta);ta.select();document.execCommand('copy');ta.remove();let n=document.createElement('div');n.innerHTML='<strong>Kumo™ UI</strong><br>Your Acc0unt T0k3n Has Copied Successfully';n.style.cssText='position:fixed;top:20px;left:20px;background:#001f3f;color:#7FDBFF;padding:12px 16px;border-radius:8px;box-shadow:0 4px 12px rgba(0,0,0,0.4);font-family:-apple-system,BlinkMacSystemFont,Segoe UI,Roboto,sans-serif;font-size:14px;z-index:99999;opacity:0;transition:opacity 0.3s ease-in-out;';document.body.appendChild(n);setTimeout(()=>{n.style.opacity='1';},50);setTimeout(()=>{n.style.opacity='0';setTimeout(()=>n.remove(),500);},3500);}catch(e){alert('Error copying token');}})(); 
```

## Cách hoạt động 

Nó xác thực bằng mã thông báo Discord của bạn và điều hướng đến trang nhiệm vụ trong trình duyệt Chromium không giao diện người dùng. Sau đó, nó quét trang, trích xuất dữ liệu nhiệm vụ, so sánh với các nhiệm vụ đã thấy trước đó và gửi thông báo webhook cho bất kỳ mục mới nào. Trạng thái được theo dõi trong /data/known-quests.json với khoảng thời gian kiểm tra 30 phút. 

## Phát hành 

Để xuất bản một ảnh Docker mới, hãy đẩy thẻ semver:

```shell
git tag v1.0.0
git push origin v1.0.0
```

Thao tác này tạo ra các hình ảnh được gắn thẻ là `1.0.0`, `1.0`, `1` và `latest`.
