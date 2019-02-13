
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Server.Models;

namespace Server.Util
{
    public class VideoService
    {
        public async Task<IEnumerable<VideoResponse>> GetAll()
        {
            int id = 0;
            VideoResponse[] result = new[]
            {
                new VideoResponse { Id = ++id, Name = $"vod{id}.mp4", Data = new byte[0] },
                new VideoResponse { Id = ++id, Name = $"vod{id}.mp4", Data = new byte[0] },
                new VideoResponse { Id = ++id, Name = $"vod{id}.mp4", Data = new byte[0] },
                new VideoResponse { Id = ++id, Name = $"vod{id}.mp4", Data = new byte[0] }
            };
            return result;

        }
        public async Task<VideoResponse> GetById(int id)
        {
            return new VideoResponse { Id = id, Name = $"vod{id}.mp4", Data = new byte[0] };
        }

        internal async Task<bool> DeleteById(int id)
        {
            return true;
        }
    }
}
