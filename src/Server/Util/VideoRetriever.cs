
using System.Collections.Generic;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Server.Models;

namespace Server.Util
{
    public class VideoRetriever
    {
        public Task<IEnumerable<VideoResponse>> GetAll()
        {
            int id = 0;
            return Task.FromResult(new[]
            {
                new VideoResponse { Id = ++id, Name = $"vod{id}.mp4", Data = new byte[0] },
                new VideoResponse { Id = ++id, Name = $"vod{id}.mp4", Data = new byte[0] },
                new VideoResponse { Id = ++id, Name = $"vod{id}.mp4", Data = new byte[0] },
                new VideoResponse { Id = ++id, Name = $"vod{id}.mp4", Data = new byte[0] }
            });

        }
        public Task<VideoResponse> GetById(int id)
        {
            return Task.FromResult(new VideoResponse { Id = id, Name = $"vod{id}.mp4", Data = new byte[0] });
        }
    }
}
