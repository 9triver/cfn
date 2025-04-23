package com.github9triver.cfn.web;

import com.github9triver.cfn.manager.LocalResourceManager;
import com.github9triver.cfn.model.dto.ResourceDto;
import com.github9triver.cfn.model.vo.ResourceVo;
import com.github9triver.cfn.model.vo.Response;
import com.github9triver.cfn.util.mapper.ResourceMapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api")
public class ResourceController {

    private LocalResourceManager localResourceManager;

    @Autowired
    public void setLocalResourceManager(LocalResourceManager localResourceManager) {
        this.localResourceManager = localResourceManager;
    }

    @GetMapping("/resources/total")
    public Response<ResourceVo> getTotalResourceCount() {
        ResourceDto resources = localResourceManager.getTotalResourceCount();
        return Response.ok(ResourceMapper.dto2vo(resources));
    }

    @GetMapping("/resources/available")
    public Response<ResourceVo> getAvailableResourceCount() {
        ResourceDto resources = localResourceManager.getAvailableResourceCount();
        return Response.ok(ResourceMapper.dto2vo(resources));
    }

}
