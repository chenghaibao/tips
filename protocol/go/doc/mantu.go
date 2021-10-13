package doc

/**
 * request:
 * | - head size - | - stream | empty  | main type | second type - | class size | class  | method size | method |
 * |     2 byte    | 3bit - 3bit empty - 6bit main - 4bit second - | 1 byte     | N byte | 1 byte      | N byte |
 *
 * | - request trace id - |
 * |     16 byte          |
 *
 * | - header size - | - header len - | .. | key size | key    | value len | .. | value size | value  | .. | .. |
 * |     1 byte      |     1 byte     |    |  1 byte  | N byte | 1 byte    |    | 2 byte     | N byte |    |    |
 *
 * ===============================================================================================================
 *
 *
 * response:
 * | - head size - | - stream | empty  | main type | second type - |
 * |     2 byte    | 3bit - 3bit empty - 6bit main - 4bit second - |
 *
 * | - header size - | - header len - | .. | key size | key    | value len | .. | value size | value  | .. | .. |
 * |     1 byte      |     1 byte     |    |  1 byte  | N byte | 1 byte    |    | 2 byte     | N byte |    |    |
 *
 * -------------- extend struct ------------------
 *
 * ==== return error ====
 * | - trace id - | - error code - | - err size - | - err msg - |
 * | 16 byte      | 8 byte         |    4 byte    |    N byte   |
 */
