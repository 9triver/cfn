package com.github9triver.cfn.model.vo;

import io.swagger.v3.oas.annotations.media.Schema;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.Getter;
import lombok.NoArgsConstructor;

import java.io.Serial;
import java.io.Serializable;

@Data
@NoArgsConstructor
public class Response<T> implements Serializable {

    @Serial
    private static final long serialVersionUID = 1L;
    protected Status status;

    protected T result;

    public Response(Status status) {
        this.status = status;
    }

    public Response(T t) {
        status = Status.newStatus(StatusEnum.SUCCESS);
        this.result = t;
    }

    public static <T> Response<T> ok(T t) {
        return new Response<T>(t);
    }

    public static <T> Response<T> fail(StatusEnum status, Object... args) {
        return new Response<>(Status.newStatus(status, args));
    }

    public static <T> Response<T> fail(Status status) {
        return new Response<>(status);
    }


    @Data
    @NoArgsConstructor
    @AllArgsConstructor
    public static class Status {

        /**
         * 业务状态码
         */
        @Schema(description = "状态码, 0 表示成功返回，其他异常返回", requiredMode = Schema.RequiredMode.REQUIRED, example = "0")
        private int code;

        /**
         * 描述信息
         */
        @Schema(description = "正确返回时为 ok，异常时为描述文案", requiredMode = Schema.RequiredMode.REQUIRED, example = "ok")
        private String msg;

        public static Status newStatus(int code, String msg) {
            return new Status(code, msg);
        }

        public static Status newStatus(StatusEnum status, Object... msgs) {
            String msg;
            if (msgs.length > 0) {
                msg = String.format(status.getMsg(), msgs);
            } else {
                msg = status.getMsg();
            }
            return newStatus(status.getCode(), msg);
        }
    }

    /**
     * 异常码规范：
     * xxx - xxx - xxx
     * 业务 - 状态 - code
     * <p>
     * 业务取值
     * - 100 全局
     * - 200 文章相关
     * - 300 评论相关
     * - 400 用户相关
     * <p>
     * 状态：基于http status的含义
     * - 4xx 调用方使用姿势问题
     * - 5xx 服务内部问题
     * <p>
     * code: 具体的业务code
     *
     * @author XuYifei
     * @date 2024-07-12
     */
    @Getter
    public enum StatusEnum {
        SUCCESS(0, "OK"),

        // -------------------------------- 通用

        // 全局传参异常
        ILLEGAL_ARGUMENTS(100_400_001, "参数异常"),
        ILLEGAL_ARGUMENTS_MIXED(100_400_002, "参数异常:%s"),

        // 全局权限相关
        FORBID_ERROR(100_403_001, "无权限"),

        FORBID_ERROR_MIXED(100_403_002, "无权限:%s"),
        FORBID_NOTLOGIN(100_403_003, "未登录"),

        // 全局，数据不存在
        RECORDS_NOT_EXISTS(100_404_001, "记录不存在:%s"),

        // 系统异常
        UNEXPECT_ERROR(100_500_001, "非预期异常:%s"),

        // 图片相关异常类型
        UPLOAD_PIC_FAILED(100_500_002, "图片上传失败！"),

        // --------------------------------

        // 文章相关异常类型，前缀为200
        ARTICLE_NOT_EXISTS(200_404_001, "文章不存在:%s"),
        COLUMN_NOT_EXISTS(200_404_002, "教程不存在:%s"),
        COLUMN_QUERY_ERROR(200_500_003, "教程查询异常:%s"),
        // 教程文章已存在
        COLUMN_ARTICLE_EXISTS(200_500_004, "专栏教程已存在:%s"),
        ARTICLE_RELATION_TUTORIAL(200_500_006, "文章已被添加为教程:%s"),
        // 分类不存在
        CATEGORY_NOT_EXISTS(200_404_101, "分类不存在:%s"),

        // --------------------------------

        // 评论相关异常类型
        COMMENT_NOT_EXISTS(300_404_001, "评论不存在:%s"),


        // --------------------------------

        // 用户相关异常
        LOGIN_FAILED_MIXED(400_403_001, "登录失败:%s"),
        USER_NOT_EXISTS(400_404_001, "用户不存在:%s"),
        USER_EXISTS(400_404_002, "用户已存在:%s"),
        // 用户登录名重复
        USER_LOGIN_NAME_REPEAT(400_404_003, "用户登录名重复:%s"),
        // 待审核
        USER_NOT_AUDIT(400_500_001, "用户未审核:%s"),
        // 星球编号不存在
        USER_STAR_NOT_EXISTS(400_404_002, "星球编号不存在:%s"),
        // 星球编号重复
        USER_STAR_REPEAT(400_404_002, "星球编号重复:%s"),
        USER_PWD_ERROR(400_500_002, "用户名or密码错误"),
        // 权限不足
        NO_PERMISSION(400_403_002, "权限不足");

        private final int code;

        private final String msg;

        StatusEnum(int code, String msg) {
            this.code = code;
            this.msg = msg;
        }

        public static boolean is5xx(int code) {
            return code % 1000_000 / 1000 >= 500;
        }

        public static boolean is403(int code) {
            return code % 1000_000 / 1000 == 403;
        }

        public static boolean is4xx(int code) {
            return code % 1000_000 / 1000 < 500;
        }
    }

}
