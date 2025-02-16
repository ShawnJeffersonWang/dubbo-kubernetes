/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import request from '@/base/http/request'

export const searchService = (params: any): Promise<any> => {
  return request({
    url: '/service/search',
    method: 'get',
    params
  })
}

export const getServiceDetail = (params: any): Promise<any> => {
  return request({
    url: '/service/detail',
    method: 'get',
    params
  })
}

export const getServiceDistribution = (params: any): Promise<any> => {
  return request({
    url: '/service/distribution',
    method: 'get',
    params
  })
}

//Get timeout time.
export const getServiceTimeout = (params: any): Promise<any> => {
  return request({
    url: '/service/config/timeout',
    method: 'get',
    params
  })
}
