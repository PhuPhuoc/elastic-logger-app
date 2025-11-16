package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseError gửi phản hồi lỗi HTTP phù hợp dựa trên kiểu lỗi.
// Nếu lỗi là *AppError, nó sử dụng các trường của nó để xây dựng phản hồi.
// Nếu lỗi là một lỗi khác, nó trả về mã 400 (Bad Request) với thông báo lỗi cơ bản.
func ResponseError(c *gin.Context, err error) {
	if apperr, ok := err.(*AppError); ok {
		// Trong môi trường không phải debug, tránh gửi thông tin lỗi nội bộ (Inner) cho client.
		if !gin.IsDebugging() {
			// Tạo một bản sao của lỗi nhưng không có thông tin Inner để bảo mật.
			// Sử dụng các trường khác như Code, Message, ReasonField, Details, File, Line, Function, Timestamp, ErrorID.
			// Tuy nhiên, việc thay đổi trực tiếp Inner như sau có thể ảnh hưởng đến đối tượng gốc nếu được sử dụng ở nơi khác.
			// Một cách an toàn hơn là tạo một đối tượng mới hoặc đảm bảo rằng `WithInner` tạo một bản sao.
			// Vì `WithInner` hiện tại thay đổi chính đối tượng `e`, ta cần cẩn thận.
			// Một cách tiếp cận an toàn hơn là gửi đối tượng gốc nhưng đảm bảo `Inner` không được serialize (đã có `json:"-"`).
			// Trong JSON gửi đi, trường `Inner` sẽ không xuất hiện.
			// Tuy nhiên, nếu bạn muốn đảm bảo giá trị `Inner` là rỗng trong mọi trường hợp, bạn cần xử lý riêng.
			// Cách đơn giản: gửi trực tiếp `apperr`, vì `json:"-"` sẽ ẩn `Inner` trong JSON.
			// Nhưng nếu bạn muốn `Inner` là "" trong cả cấu trúc được gửi (nếu cấu trúc được dùng ở nơi khác sau khi gọi hàm này),
			// thì bạn cần tạo một bản sao. Tuy nhiên, việc tạo bản sao toàn bộ *AppError chỉ để thay đổi `Inner` là phức tạp.
			// Cách tiếp cận hiện tại với `json:"-"` là chuẩn và an toàn cho JSON response.
			// Nếu bạn lo lắng về việc `apperr` được sử dụng sau này với `Inner` còn đó, thì `WithInner` hiện tại thay đổi chính nó,
			// nên nếu gọi `apperr.WithInner("")` trước khi trả về, nó sẽ thay đổi `apperr` gốc.
			// Để không thay đổi `apperr` gốc, ta nên tạo một bản sao chỉ với Inner là rỗng.
			// Nhưng để đơn giản hóa, cách dùng `json:"-"` là cách phổ biến nhất.
			// Nếu bạn vẫn muốn chắc chắn rằng Inner không có giá trị trong cấu trúc được trả về trong chế độ không debug,
			// bạn có thể cần phải tạo một struct mới hoặc xử lý phức tạp hơn.
			// Tuy nhiên, với `json:"-"`, phản hồi JSON sẽ không có trường Inner, điều này là đủ cho mục đích bảo mật.
			// Vì vậy, gửi trực tiếp `apperr` là hợp lý, vì `Inner` sẽ không xuất hiện trong JSON.
			// Câu lệnh sau: `errWithNoInner := apperr.WithInner(nil)` không hiệu quả vì `WithInner` trả về `e` (chính nó).
			// `apperr.WithInner(nil)` sẽ làm `apperr.Inner = nil`, ảnh hưởng đến đối tượng gốc.
			// Cách tốt nhất là giữ nguyên và tin tưởng `json:"-"`.
			// Tuy nhiên, để minh họa cách làm rõ ràng hơn nếu cần thay đổi giá trị Inner trong phản hồi mà không ảnh hưởng gốc,
			// ta có thể tạo một bản sao tạm thời *nếu cần*. Nhưng `json:"-"` đã xử lý việc ẩn trường này trong JSON.
			// Cân nhắc lại: `WithInner("")` hoặc `WithInner(nil)` sẽ thay đổi `apperr` gốc. Không nên làm vậy.
			// Cách tốt nhất là gửi `apperr` trực tiếp, vì `json:"-"` đảm bảo `Inner` không xuất hiện trong JSON.
			// Tuy nhiên, ví dụ trước đó đã dùng `apperr.WithInner("")`, điều này thay đổi đối tượng gốc.
			// Để không thay đổi đối tượng gốc, ta có thể không gọi `WithInner` nữa và tin tưởng `json:"-"`.
			// Nhưng nếu bạn muốn giữ nguyên logic gọi một phương thức, bạn có thể tạo một phương thức mới
			// trả về một *AppError mới mà không có Inner, hoặc đảm bảo `WithInner` tạo bản sao.
			// Ví dụ, `WithInner` có thể được sửa để tạo bản sao, nhưng điều đó thay đổi ý nghĩa của nó.
			// Cách đơn giản và hiệu quả nhất là:
			// 1. Giữ `Inner` với `json:"-"`.
			// 2. Trong chế độ debug, gửi `apperr` nguyên bản.
			// 3. Trong chế độ không debug, gửi `apperr`, trường `Inner` sẽ tự động bị ẩn khỏi JSON nhờ `json:"-"`.
			// Do đó, code sau là phù hợp và an toàn:
			c.JSON(apperr.StatusCode(), gin.H{
				"success": false,
				"error":   apperr, // Trường Inner sẽ không xuất hiện trong JSON do `json:"-"`.
			})
			return
		}

		// Trong môi trường debug, gửi toàn bộ thông tin lỗi, bao gồm cả lỗi nội bộ (Inner).
		c.JSON(apperr.StatusCode(), gin.H{
			"success": false,
			"error":   apperr,
		})
		return
	}

	// Nếu lỗi không phải là *AppError, trả về lỗi chung với mã 400.
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"status":  http.StatusBadRequest,
		"error": gin.H{
			"message": "An unexpected error occurred.",
			"reason":  "The error type is not a structured AppError.",
		},
	})
}
