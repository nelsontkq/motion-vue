
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;

namespace Server.Models
{
    public class VideoResponse
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public byte[] Data { get; set; }
    }
}
